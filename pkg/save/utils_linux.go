//go:build linux

package save

import (
	"os"
	"path/filepath"
)

func DefaultSaveLocation() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, "snap/steam/common/.local/share/Steam/steamapps/compatdata/1511460/pfx/drive_c/users/steamuser/")
	path = saveFolderWindows(path)
	if _, err := os.Stat(path); path != "" && !os.IsNotExist(err) {
		return path, nil
	} else {
		return home, nil
	}
}
