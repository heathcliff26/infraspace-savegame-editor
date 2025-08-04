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
	oldGitCommit := gitCommit
	defer func() { gitCommit = oldGitCommit }()

	gitCommit = "1234567890abcdef"

	a := test.NewApp()
	v := NewVersionFromApp(a)

	assert := assert.New(t)

	if a.Metadata().Name != "" {
		assert.Contains(a.Metadata().Name, v.Name)
	} else {
		assert.Contains(os.Args[0], v.Name)
	}
	assert.Equal("v"+a.Metadata().Version, v.Version)
	assert.Equal("1234567", v.Commit, "commit hash should be truncated")
	assert.Equal(runtime.Version(), v.Go)
	assert.Equal(save.GAME_VERSION, v.GameVersion)
}

func TestInitGitCommit(t *testing.T) {
	oldGitCommit := gitCommit
	defer func() { gitCommit = oldGitCommit }()
	assert := assert.New(t)

	gitCommit = "1234567890abcdef"
	initGitCommit()
	assert.Equal("1234567890abcdef", gitCommit, "gitCommit should not be changed")

	gitCommit = "$Format:%H$"
	initGitCommit()
	assert.NotEqual("$Format:%H$", gitCommit, "gitCommit should be changed")
}
