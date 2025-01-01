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

	paths := []string{
		".local/share/Steam/steamapps/compatdata/1511460/pfx/drive_c/users/steamuser/",
		"snap/steam/common/.local/share/Steam/steamapps/compatdata/1511460/pfx/drive_c/users/steamuser/",
		".var/app/com.valvesoftware.Steam/data/Steam/steamapps/compatdata/1511460/pfx/drive_c/users/steamuser/",
	}

	for _, path := range paths {
		path = filepath.Join(home, path)
		path = saveFolderWindows(path)
		if _, err := os.Stat(path); path != "" && !os.IsNotExist(err) {
			return path, nil
		}
	}
	return home, nil
}
