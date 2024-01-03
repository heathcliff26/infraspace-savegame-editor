package save

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func readSaveFile(path string) (string, []byte, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return "", nil, err
	}
	fullSave := string(buf)
	i := strings.Index(fullSave, "{")
	if i < 0 {
		return "", nil, fmt.Errorf("could not find start of json-body")
	}
	return fullSave[:i], []byte(fullSave[i:]), nil
}

func maxBuildingStorage(building Building) Building {
	if building.ConsumerProducer == nil {
		return building
	}
	for i := range building.ConsumerProducer.IncomingStorage {
		building.ConsumerProducer.IncomingStorage[i] = BUILDING_MAX_STORAGE
	}
	return building
}

func marshalJSON(v any) (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)

	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")

	err := enc.Encode(v)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
