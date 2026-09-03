//go:build linux

package save

import (
	"os"
	"path/filepath"

	"github.com/andygrunwald/vdf"
)

func DefaultSaveLocation() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	libraryfolders := loadLibraryFolders(home)
	if libraryfolders == nil {
		return home, nil
	}

	libraryfolders = libraryfolders["libraryfolders"].(map[string]interface{})

	for _, v := range libraryfolders {
		path := v.(map[string]interface{})["path"].(string)
		path = filepath.Join(path, "steamapps/compatdata/1511460/pfx/drive_c/users/steamuser/")
		path = saveFolderWindows(path)
		if _, err := os.Stat(path); path != "" && !os.IsNotExist(err) {
			return path, nil
		}
	}

	return home, nil
}

func loadLibraryFolders(home string) map[string]interface{} {
	paths := []string{
		".local/share/Steam/config/libraryfolders.vdf",
		".var/app/com.valvesoftware.Steam/data/Steam/config/libraryfolders.vdf",
		"snap/steam/common/.local/share/Steam/config/libraryfolders.vdf",
	}

	for _, path := range paths {
		path = filepath.Join(home, path)
		f, err := os.Open(path)
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return nil
		}
		defer f.Close()

		p := vdf.NewParser(f)
		cfg, err := p.Parse()
		if err != nil {
			return nil
		}
		return cfg
	}
	return nil
}
