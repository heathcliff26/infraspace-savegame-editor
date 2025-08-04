package gui

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/heathcliff26/infraspace-savegame-editor/pkg/save"
)

const RELEASE = true

// NOTE: The $Format strings are replaced during 'git archive' thanks to the
// companion .gitattributes file containing 'export-subst' in this same
// directory.  See also https://git-scm.com/docs/gitattributes
var gitCommit string = "$Format:%H$" // sha1 from git, output of $(git rev-parse HEAD)

func init() {
	initGitCommit()
}

func initGitCommit() {
	if strings.HasPrefix(gitCommit, "$Format") {
		var commit string
		buildinfo, _ := debug.ReadBuildInfo()
		for _, item := range buildinfo.Settings {
			if item.Key == "vcs.revision" {
				commit = item.Value
				break
			}
		}
		gitCommit = commit
	}
}

type Version struct {
	Name, Version, Commit, Go, GameVersion string
}

func NewVersionFromApp(app fyne.App) Version {
	commit := gitCommit
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
	data := [][]string{
		{"Version:", v.Version},
		{"Commit:", v.Commit},
		{"Go:", v.Go},
		{"InfraSpace:", v.GameVersion},
	}

	versionTable := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("                    ")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)

	versionTable.ShowHeaderRow = false
	versionTable.ShowHeaderColumn = false
	versionTable.StickyRowCount = len(data) - 1
	versionTable.StickyColumnCount = len(data[0]) - 1
	versionTable.HideSeparators = true

	return versionTable
}
