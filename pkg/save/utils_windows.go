//go:build windows

package save

import (
	"os"
)

func DefaultSaveLocation() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := saveFolderWindows(home)
	if _, err := os.Stat(path); path != "" && !os.IsNotExist(err) {
		return path, nil
	} else {
		return home, nil
	}
}
