package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type savegame struct {
	path   string
	prefix string
	data   map[string]interface{}
}

// Load the savegame from the path
func LoadSavegame(path string) (*savegame, error) {
	buf, err := ioutil.ReadFile(path)
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
	err = ioutil.WriteFile(save.getPath(), []byte(fullSave), 0644)
	return err
}

func (save *savegame) getPath() string {
	return save.path
}

func (save *savegame) setPath(newPath string) {
	save.path = newPath
}

func (save *savegame) getPrefix() string {
	return save.prefix
}

func (save *savegame) setPrefix(newPrefix string) {
	save.prefix = newPrefix
}

// Points to save data as a JSON Object (map[string]interface{})
func (save *savegame) Data() map[string]interface{} {
	return save.data
}

func (save *savegame) getResearchProgress() map[string]interface{} {
	return save.Data()["researchManager"].(map[string]interface{})["researchProgress"].(map[string]interface{})
}

func (save *savegame) getStarterWorkers() []interface{} {
	return save.Data()["market"].(map[string]interface{})["starterWorkers"].([]interface{})
}

func (save *savegame) getResources() map[string]interface{} {
	return save.Data()["resources"].(map[string]interface{})
}

func (save *savegame) getBuildings() []interface{} {
	return save.Data()["buildings"].([]interface{})
}
