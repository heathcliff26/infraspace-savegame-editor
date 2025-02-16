package gui

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/heathcliff26/infraspace-savegame-editor/pkg/save"
)

const RELEASE = true

type Version struct {
	Name, Version, Commit, Go, GameVersion string
}

func NewVersionFromApp(app fyne.App) Version {
	var commit string
	buildinfo, _ := debug.ReadBuildInfo()
	for _, item := range buildinfo.Settings {
		if item.Key == "vcs.revision" {
			commit = item.Value
			break
		}
	}
	if len(commit) > 7 {
		commit = commit[:7]
	}

	metadata := app.Metadata()

	name, _ := strings.CutSuffix(metadata.Name, ".exe")
	if name == "" {
		name = filepath.Base(os.Args[0])
	}

	result := Version{
		Name:        name,
		Version:     "v" + metadata.Version,
		Commit:      commit,
		Go:          runtime.Version(),
		GameVersion: save.GAME_VERSION,
	}

	if !RELEASE {
		result.Version += "-devel"
	}

	return result
}

// Create the content for the version dialog
func (v Version) CreateContent() fyne.CanvasObject {
	r1 := make([]fyne.CanvasObject, 4)
	r2 := make([]fyne.CanvasObject, 4)
	r1[0] = canvas.NewText("Version:", TEXT_COLOR)
	r2[0] = canvas.NewText(v.Version, TEXT_COLOR)
	r1[1] = canvas.NewText("Commit:", TEXT_COLOR)
	r2[1] = canvas.NewText(v.Commit, TEXT_COLOR)
	r1[2] = canvas.NewText("Go:", TEXT_COLOR)
	r2[2] = canvas.NewText(v.Go, TEXT_COLOR)
	r1[3] = canvas.NewText("InfraSpace:", TEXT_COLOR)
	r2[3] = canvas.NewText(v.GameVersion, TEXT_COLOR)

	row1 := container.NewVBox(r1...)
	row2 := container.NewVBox(r2...)

	return container.NewPadded(container.NewHBox(row1, row2))
}
