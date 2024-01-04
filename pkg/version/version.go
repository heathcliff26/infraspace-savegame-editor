package version

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"fyne.io/fyne/v2/app"
	"github.com/heathcliff26/infraspace-savegame-editor/pkg/save"
)

type version struct {
	Name, Version, Commit, Go, GameVersion string
}

func Version() version {
	var commit string
	buildinfo, _ := debug.ReadBuildInfo()
	for _, item := range buildinfo.Settings {
		if item.Key == "vcs.revision" {
			commit = item.Value
			break
		}
	}

	metadata := app.New().Metadata()

	name, _ := strings.CutSuffix(metadata.Name, ".exe")
	if name == "" {
		name = filepath.Base(os.Args[0])
	}

	result := version{
		Name:        name,
		Version:     metadata.Version,
		Commit:      commit,
		Go:          runtime.Version(),
		GameVersion: save.GAME_VERSION,
	}

	if !metadata.Release {
		result.Version += "-devel"
	}

	return result
}

func (v version) String() string {
	result := v.Name + "\n"
	result += "Version:    " + v.Version + "\n"
	result += "Commit:     " + v.Commit + "\n"
	result += "Go:         " + v.Go + "\n"
	result += "InfraSpace: " + v.GameVersion + "\n"
	return result
}
