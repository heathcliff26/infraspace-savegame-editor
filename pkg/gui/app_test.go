package gui

import (
	"image/color"
	"testing"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
	"github.com/heathcliff26/infraspace-savegame-editor/pkg/save"
	"github.com/stretchr/testify/assert"
)

func TestNewGUI(t *testing.T) {
	g := &GUI{
		App:  test.NewApp(),
		Main: test.NewWindow(canvas.NewText("Test", color.White)),
	}

	t.Run("Menu", func(t *testing.T) {
		g.makeMenu()

		assert.NotEmpty(t, g.Menu)
	})
	t.Run("Resources", func(t *testing.T) {
		g.makeResourcesBox()

		assert := assert.New(t)

		if !assert.NotEmpty(g.Resources) {
			t.FailNow()
		}

		names := make([]string, len(g.Resources))
		for i, resource := range g.Resources {
			names[i] = resource.Name
			v, err := resource.Value.Get()
			assert.Nil(err)
			assert.Zero(v)
			assert.NotNil(resource.Entry)
		}
		assert.ElementsMatch(save.ResourceNames(), names)
	})
	t.Run("Research", func(t *testing.T) {
		g.makeResearchBox()

		assert := assert.New(t)

		if !assert.NotEmpty(t, g.Research) {
			t.FailNow()
		}

		if !assert.NotNil(g.UnlockAllResearch) {
			t.FailNow()
		}

		assert.False(g.UnlockAllResearch.Checked)

		for _, tValue := range []bool{true, false} {
			g.UnlockAllResearch.OnChanged(tValue)
			for _, research := range g.Research {
				assert.Equal(tValue, research.Checkbox.Disabled())
			}
		}

		names := make([]string, len(g.Research))
		for i, research := range g.Research {
			names[i] = research.Name
			if !assert.NotNil(research.Checkbox) {
				t.FailNow()
			}
			assert.False(research.Checkbox.Checked)
		}
		assert.ElementsMatch(save.ResearchNames(), names)
	})
	t.Run("SpaceshipParts", func(t *testing.T) {
		g.makeSpaceshipBox()

		assert := assert.New(t)

		if !assert.NotEmpty(g.SpaceshipParts) {
			t.FailNow()
		}

		names := make([]string, len(g.SpaceshipParts))
		for i, part := range g.SpaceshipParts {
			names[i] = part.Name
			if !assert.NotNil(part.Checkbox) {
				t.FailNow()
			}
			assert.False(part.Checkbox.Checked)
		}
		assert.ElementsMatch(save.SpaceshipParts(), names)
	})
	t.Run("OtherOptions", func(t *testing.T) {
		g.makeOptionsBox()
		// TODO
		assert.NotEmpty(t, g.OtherOptions)
	})
	t.Run("ActionButtons", func(t *testing.T) {
		g.makeActionButtons()

		assert := assert.New(t)

		if !assert.NotEmpty(g.ActionButtons) {
			t.FailNow()
		}

		for _, b := range g.ActionButtons {
			assert.True(b.Disabled())
		}
	})
}
