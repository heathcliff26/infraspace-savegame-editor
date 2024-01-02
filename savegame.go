package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type savegame struct {
	path   string
	prefix string
	data   map[string]interface{}
}

// Load the savegame from the path
func LoadSavegame(path string) (*savegame, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fullSave := string(buf)
	i := strings.Index(fullSave, "{")
	if i < 0 {
		return nil, fmt.Errorf("Could not find start of JSON")
	}
	prefix := fullSave[:i]
	buf = []byte(fullSave[i:])
	var j interface{}
	err = json.Unmarshal(buf, &j)
	if err != nil {
		return nil, err
	}
	data := j.(map[string]interface{})
	return &savegame{
		path:   path,
		prefix: prefix,
		data:   data,
	}, nil
}

// Save the savegame to the Path
func (save *savegame) Save() error {
	buf, err := json.MarshalIndent(save.Data(), "", "  ")
	if err != nil {
		return err
	}
	fullSave := save.getPrefix() + string(buf)
	err = os.WriteFile(save.getPath(), []byte(fullSave), 0644)
	return err
}

func (save *savegame) getPath() string {
	return save.path
}

// func (save *savegame) setPath(newPath string) {
// 	save.path = newPath
// }

func (save *savegame) getPrefix() string {
	return save.prefix
}

// func (save *savegame) setPrefix(newPrefix string) {
// 	save.prefix = newPrefix
// }

// Points to save data as a JSON Object (map[string]interface{})
func (save *savegame) Data() map[string]interface{} {
	return save.data
}

func (save *savegame) getResearchProgress() map[string]interface{} {
	result, _ := GetJSONObject(save.Data(), "researchManager", "researchProgress")
	return result
}

func (save *savegame) getStarterWorkers() []interface{} {
	result, _ := GetJSONArray(save.Data(), "market", "starterWorkers")
	return result
}

func (save *savegame) getResources() map[string]interface{} {
	result, _ := GetJSONObject(save.Data(), "resources")
	return result
}

func (save *savegame) getBuildings() []interface{} {
	result, _ := GetJSONArray(save.Data(), "buildings")
	return result
}

func (save *savegame) getNextID() int {
	return int(save.Data()["nextID"].(float64))
}

func (save *savegame) setNextID(nextID int) {
	save.Data()["nextID"] = float64(nextID)
}
