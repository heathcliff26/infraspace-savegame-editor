package save

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ResourceNames() []string {
	list := make([]string, len(resourceNames))
	copy(list, resourceNames)
	return list
}

func ResearchNames() []string {
	list := make([]string, len(researchNames))
	copy(list, researchNames)
	return list
}

func SpaceshipParts() []string {
	list := make([]string, len(spaceshipParts))
	copy(list, spaceshipParts)
	return list
}

func saveFolderWindows(root string) string {
	return filepath.Join(root, "AppData", "LocalLow", "Dionic Software", "InfraSpace", "saves")
}

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

func maxFactoryStorage(building Building) Building {
	if building.ConsumerProducer == nil {
		return building
	}
	building.ConsumerProducer.IncomingStorage = maxStorage(building.ConsumerProducer.IncomingStorage)
	if building.BuildingName != "spaceShipConstructionFacility" {
		building.ConsumerProducer.OutgoingStorage = maxStorage(building.ConsumerProducer.OutgoingStorage)
	}

	return building
}

func maxStorage(storage []int64) []int64 {
	for i := range storage {
		storage[i] = BUILDING_MAX_STORAGE
	}
	return storage
}

func maxHabitatStorage(building Building) Building {
	if building.ConsumerProducer == nil || building.ConsumerProducer.Type != TYPE_HABITAT {
		return building
	}
	for key := range building.ConsumerProducer.ProductionLogic.(HabitatProductionLogic).Storage {
		building.ConsumerProducer.ProductionLogic.(HabitatProductionLogic).Storage[key] = json.Number(strconv.Itoa(BUILDING_MAX_STORAGE))
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
