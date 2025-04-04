package cache

import (
	"os"
	"sync/atomic"
	"time"

	"fyne.io/fyne/v2"
)

var (
	cacheDuration     = 1 * time.Minute
	cleanTaskInterval = cacheDuration / 2

	lastClean                     time.Time
	skippedCleanWithCanvasRefresh = false

	// testing purpose only
	timeNow = time.Now
)

func init() {
	if t, err := time.ParseDuration(os.Getenv("FYNE_CACHE")); err == nil {
		cacheDuration = t
		cleanTaskInterval = cacheDuration / 2
	}
}

// Clean run cache clean task, it should be called on paint events.
func Clean(canvasRefreshed bool) {
	now := timeNow()
	// do not run clean task too fast
	if now.Sub(lastClean) < 10*time.Second {
		if canvasRefreshed {
			skippedCleanWithCanvasRefresh = true
		}
		return
	}
	if skippedCleanWithCanvasRefresh {
		skippedCleanWithCanvasRefresh = false
		canvasRefreshed = true
	}
	if !canvasRefreshed && now.Sub(lastClean) < cleanTaskInterval {
		return
	}
	destroyExpiredSvgs(now)
	destroyExpiredFontMetrics(now)
	if canvasRefreshed {
		// Destroy renderers on canvas refresh to avoid flickering screen.
		destroyExpiredRenderers(now)
		// canvases cache should be invalidated only on canvas refresh, otherwise there wouldn't
		// be a way to recover them later
		destroyExpiredCanvases(now)
	}
	lastClean = timeNow()
}

// CleanCanvas performs a complete remove of all the objects that belong to the specified
// canvas. Usually used to free all objects from a closing windows.
func CleanCanvas(canvas fyne.Canvas) {
	deletingObjs := make([]fyne.CanvasObject, 0, 50)

	canvasesLock.RLock()
	for obj, cinfo := range canvases {
		if cinfo.canvas == canvas {
			deletingObjs = append(deletingObjs, obj)
		}
	}
	canvasesLock.RUnlock()
	if len(deletingObjs) == 0 {
		return
	}

	canvasesLock.Lock()
	for _, dobj := range deletingObjs {
		delete(canvases, dobj)
	}
	canvasesLock.Unlock()

	renderersLock.Lock()
	for _, dobj := range deletingObjs {
		wid, ok := dobj.(fyne.Widget)
		if !ok {
			continue
		}
		rinfo, ok := renderers[wid]
		if !ok {
			continue
		}
		rinfo.renderer.Destroy()
		overrides.Delete(wid)
		delete(renderers, wid)
	}
	renderersLock.Unlock()
}

// CleanCanvases runs cache clean tasks for canvases that are being refreshed. This is called on paint events.
func CleanCanvases(refreshingCanvases []fyne.Canvas) {
	now := timeNow()

	// do not run clean task too fast
	if now.Sub(lastClean) < 10*time.Second {
		return
	}

	if now.Sub(lastClean) < cleanTaskInterval {
		return
	}

	destroyExpiredSvgs(now)
	destroyExpiredFontMetrics(now)

	deletingObjs := make([]fyne.CanvasObject, 0, 50)

	canvasesLock.RLock()
	for obj, cinfo := range canvases {
		if cinfo.isExpired(now) && matchesACanvas(cinfo, refreshingCanvases) {
			deletingObjs = append(deletingObjs, obj)
		}
	}
	canvasesLock.RUnlock()
	if len(deletingObjs) == 0 {
		return
	}

	canvasesLock.Lock()
	for _, dobj := range deletingObjs {
		delete(canvases, dobj)
	}
	canvasesLock.Unlock()

	renderersLock.Lock()
	for _, dobj := range deletingObjs {
		wid, ok := dobj.(fyne.Widget)
		if !ok {
			continue
		}
		rinfo, ok := renderers[wid]
		if !ok || !rinfo.isExpired(now) {
			continue
		}
		rinfo.renderer.Destroy()
		overrides.Delete(wid)
		delete(renderers, wid)
	}
	renderersLock.Unlock()
	lastClean = timeNow()
}

// ResetThemeCaches clears all the svg and text size cache maps
func ResetThemeCaches() {
	svgs.Range(func(key, value any) bool {
		svgs.Delete(key)
		return true
	})

	fontSizeLock.Lock()
	fontSizeCache = map[fontSizeEntry]*fontMetric{}
	fontSizeLock.Unlock()
}

// destroyExpiredCanvases deletes objects from the canvases cache.
func destroyExpiredCanvases(now time.Time) {
	canvasesLock.Lock()
	for obj, cinfo := range canvases {
		if cinfo.isExpired(now) {
			delete(canvases, obj)
		}
	}
	canvasesLock.Unlock()
}

// destroyExpiredRenderers deletes the renderer from the cache and calls
// renderer.Destroy()
func destroyExpiredRenderers(now time.Time) {
	renderersLock.Lock()
	for wid, rinfo := range renderers {
		if rinfo.isExpired(now) {
			rinfo.renderer.Destroy()
			overrides.Delete(wid)
			delete(renderers, wid)
		}
	}
	renderersLock.Unlock()
}

// matchesACanvas returns true if the canvas represented by the canvasInfo object matches one of
// the canvases passed in 'canvases', otherwise false is returned.
func matchesACanvas(cinfo *canvasInfo, canvases []fyne.Canvas) bool {
	canvas := cinfo.canvas

	for _, obj := range canvases {
		if obj == canvas {
			return true
		}
	}
	return false
}

type expiringCache struct {
	expires atomic.Pointer[time.Time]
}

// isExpired check if the cache data is expired.
func (c *expiringCache) isExpired(now time.Time) bool {
	t := c.expires.Load()
	if t == nil {
		return (time.Time{}).Before(now)
	}
	return (*t).Before(now)
}

// setAlive updates expiration time.
func (c *expiringCache) setAlive() {
	time := timeNow().Add(cacheDuration)
	c.expires.Store(&time)
}

type expiringCacheNoLock struct {
	expires time.Time
}

// isExpired check if the cache data is expired.
func (c *expiringCacheNoLock) isExpired(now time.Time) bool {
	return c.expires.Before(now)
}

// setAlive updates expiration time.
func (c *expiringCacheNoLock) setAlive() {
	t := timeNow().Add(cacheDuration)
	c.expires = t
}
