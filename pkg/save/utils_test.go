package save

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceNames(t *testing.T) {
	res := ResourceNames()

	assert := assert.New(t)

	assert.Equal(resourceNames, res)
	assert.NotSame(&resourceNames[0], &res[0])
}

func TestResearchNames(t *testing.T) {
	res := ResearchNames()

	assert := assert.New(t)

	assert.Equal(researchNames, res)
	assert.NotSame(&researchNames[0], &res[0])
}

func TestSpaceshipParts(t *testing.T) {
	res := SpaceshipParts()

	assert := assert.New(t)

	assert.Equal(spaceshipParts, res)
	assert.NotSame(&spaceshipParts[0], &res[0])
}

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
		{
			Name:   "1.35.426",
			Prefix: "\ufeffInfraSpace\nInfraSpace 1.35.426\nCampaignMap\n01/14/2024 12:06:20\n",
		},
		{
			Name:   "1.49.446",
			Prefix: "\ufeffInfraSpace\nInfraSpace 1.49.446\nCampaignMap\n11/01/2024 09:06:42\n",
		},
		{
			Name:   "1.50.448",
			Prefix: "\ufeffInfraSpace\nInfraSpace 1.50.448\nCampaignMap\n02/16/2025 07:18:05\n",
		},
		{
			Name:   "1.53.451",
			Prefix: "\ufeffInfraSpace\nInfraSpace 1.53.451\nCampaignMap\n08/01/2025 14:49:30\n",
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			prefix, _, err := readSaveFile(filepath.Join("testdata", tCase.Name+".sav"))

			assert := assert.New(t)

			assert.Nil(err)

			expectedPrefix := tCase.Prefix
			if runtime.GOOS == "windows" {
				expectedPrefix = strings.ReplaceAll(expectedPrefix, "\n", "\r\n")
			}
			assert.Equal(expectedPrefix, prefix)
		})
	}
}

func TestMaxFactoryStorage(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Building{}, maxFactoryStorage(Building{}), "Should not change empty struct")

	b := Building{
		ConsumerProducer: &ConsumerProducer{
			IncomingStorage: []int64{0, 0, 4, 10},
			OutgoingStorage: []int64{0, 3, 7, 10},
		},
	}

	res := maxFactoryStorage(b)

	assert.Equal(b, res, "Should not be changed since ConsumerProducer is a pointer")

	expectedStorage := []int64{BUILDING_MAX_STORAGE, BUILDING_MAX_STORAGE, BUILDING_MAX_STORAGE, BUILDING_MAX_STORAGE}
	assert.Equal(expectedStorage, res.ConsumerProducer.IncomingStorage)
	assert.Equal(expectedStorage, res.ConsumerProducer.OutgoingStorage)
}

func TestMaxHabitatStorage(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Building{}, maxHabitatStorage(Building{}), "Should not change empty struct")

	b := Building{
		ConsumerProducer: &ConsumerProducer{
			Type: TYPE_HABITAT,
			ProductionLogic: HabitatProductionLogic{
				Storage: map[string]json.Number{
					"culturePoints": "0",
					"oxygen":        "0",
					"parkPoints":    "0",
					"schoolPoints":  "0",
					"survivalFood":  "0",
					"water":         "0",
				},
			},
		},
	}

	res := maxHabitatStorage(b)

	assert.Equal(b, res, "Should not be changed since ConsumerProducer is a pointer")

	buildingMaxStorage := strconv.Itoa(BUILDING_MAX_STORAGE)
	expectedStorage := map[string]json.Number{
		"culturePoints": json.Number(buildingMaxStorage),
		"oxygen":        json.Number(buildingMaxStorage),
		"parkPoints":    json.Number(buildingMaxStorage),
		"schoolPoints":  json.Number(buildingMaxStorage),
		"survivalFood":  json.Number(buildingMaxStorage),
		"water":         json.Number(buildingMaxStorage),
	}
	assert.Equal(expectedStorage, res.ConsumerProducer.ProductionLogic.(HabitatProductionLogic).Storage)
}

func TestMarshalJSON(t *testing.T) {
	input := "{\n"
	input += "  \"<disabledDueToMods>k__BackingField\": 0\n"
	input += "}\n"

	output := "{\r\n"
	output += "  \"<disabledDueToMods>k__BackingField\": 0\r\n"
	output += "}\r\n"

	assert := assert.New(t)

	var data json.RawMessage

	err := json.Unmarshal([]byte(input), &data)
	assert.Nil(err)
	res, err := marshalJSON(data)
	assert.Nil(err)
	assert.Equal(output, res)
}
