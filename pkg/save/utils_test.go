package save

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
