package save

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultSaveLocation(t *testing.T) {
	path, err := DefaultSaveLocation()
	if err != nil {
		t.Fatalf("Finished with error: %v", err)
	}

	if path == "" {
		t.Fatalf("Path should not be empty")
	}

	f, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		t.Fatalf("Does not seem to be a valid path: %v", err)
	}
	if err == nil {
		if !f.IsDir() {
			t.Fatalf("The path should be a directory")
		}
	}
}

func TestReadSaveFile(t *testing.T) {
	tMatrix := []struct {
		Name, Prefix string
	}{
		{
			Name:   "1.33.423",
			Prefix: "\ufeffInfraSpace\nInfraSpace 1.33.423\nCampaignMap\n01/03/2024 11:54:08\n",
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			prefix, _, err := readSaveFile(filepath.Join("testdata", tCase.Name+".sav"))

			assert := assert.New(t)

			assert.Nil(err)
			assert.Equal(tCase.Prefix, prefix)
		})
	}
}
