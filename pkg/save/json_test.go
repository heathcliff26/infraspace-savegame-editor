package save

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoDataLoss(t *testing.T) {
	tMatrix := []string{"1.33.423", "1.35.425", "1.35.426", "1.49.446"}

	for _, tCase := range tMatrix {
		t.Run(tCase, func(t *testing.T) {
			_, buf, err := readSaveFile(filepath.Join("testdata", tCase+".sav"))
			if err != nil {
				t.Fatalf("Failed to read save file: %v", err)
			}

			var data SaveData
			err = json.Unmarshal(buf, &data)
			if err != nil {
				t.Fatalf("Failed to unmarshal SaveData: %v", err)
			}

			parsedData, err := marshalJSON(data)
			if err != nil {
				t.Fatalf("Failed to marshal SaveData: %v", err)
			}

			assert.JSONEq(t, string(buf), parsedData)
		})
	}
}
