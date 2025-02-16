package gui

import (
	"os"
	"runtime"
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/heathcliff26/infraspace-savegame-editor/pkg/save"
	"github.com/stretchr/testify/assert"
)

func TestNewVersionFromApp(t *testing.T) {
	a := test.NewApp()
	v := NewVersionFromApp(a)

	assert := assert.New(t)

	if a.Metadata().Name != "" {
		assert.Contains(a.Metadata().Name, v.Name)
	} else {
		assert.Contains(os.Args[0], v.Name)
	}
	assert.Equal("v"+a.Metadata().Version, v.Version)
	assert.LessOrEqual(len(v.Commit), 7)
	assert.Equal(runtime.Version(), v.Go)
	assert.Equal(save.GAME_VERSION, v.GameVersion)
}
