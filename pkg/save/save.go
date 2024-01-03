package save

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Savegame struct {
	Changed bool

	path   string
	prefix string
	data   *SaveData
}

// Load the savegame from the path
func LoadSavegame(path string) (*Savegame, error) {
	prefix, buf, err := readSaveFile(path)
	if err != nil {
		return nil, err
	}

	var data SaveData
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}

	return &Savegame{
		Changed: false,
		path:    path,
		prefix:  prefix,
		data:    &data,
	}, nil
}

func (s *Savegame) Path() string {
	return s.path
}

func (s *Savegame) Prefix() string {
	return s.prefix
}

func (s *Savegame) Data() *SaveData {
	return s.data
}

// Save the savegame to the Path
func (s *Savegame) Save() error {
	data, err := marshalJSON(s.Data())
	if err != nil {
		return err
	}

	fullSave := s.prefix + data
	err = os.WriteFile(s.Path(), []byte(fullSave), 0644)
	return err
}

// Create a backup of the save file, returns the path of the backup
func (s *Savegame) Backup() (string, error) {
	dst := s.Path() + ".backup"
	i := 0
	_, err := os.Stat(dst)
	for !errors.Is(err, os.ErrNotExist) {
		i++
		dst = s.Path() + ".backup_" + fmt.Sprint(i)
		_, err = os.Stat(dst)
	}

	input, err := os.ReadFile(s.Path())
	if err != nil {
		return "", err
	}

	err = os.WriteFile(dst, input, 0644)
	if err != nil {
		return "", err
	}

	return dst, nil
}

// Print the relevant information about the savegame
func (s *Savegame) Print() {
	fmt.Println("Noop Savegame.Print()")
}

// Unlocks all research
func (s *Savegame) UnlockAllResearch() {
	for key := range s.Data().ResearchManager.ResearchProgress {
		s.Data().ResearchManager.ResearchProgress[key] = maxResearchProgress[key]
	}
	s.Changed = true
}

// Increase the starter workers ot the given count, return resulting number of starting workers
func (s *Savegame) AddStarterWorkers(count int) int64 {
	oldNextID := s.Data().NextID
	for len(s.Data().Market.StarterWorkers) < count {
		newWorker := Worker{
			Home: 0,
			ID:   s.Data().NextID,
		}
		s.Data().NextID++
		s.Data().Market.StarterWorkers = append(s.Data().Market.StarterWorkers, newWorker)
	}

	s.Changed = true
	return (s.Data().NextID - oldNextID)
}

// Get a resource by name
func (s *Savegame) GetResource(key string) int64 {
	return s.Data().Resources[key] / RESOURCE_FACTOR
}

// Set the given resource
func (s *Savegame) SetResource(key string, value int) error {
	if _, ok := s.Data().Resources[key]; ok {
		s.Data().Resources[key] = int64(value) * RESOURCE_FACTOR
		s.Changed = true
		return nil
	} else {
		return NewErrMissingKey("resources", key)
	}
}

type EditBuildingsOptions struct {
	HabitatWorkers   bool
	HabitatStorage   bool
	IndustrialRobots bool
	FactoryStorage   bool
}

// Edit the buildings with the given configuration
func (s *Savegame) EditBuildings(opt EditBuildingsOptions) {
	if opt.HabitatWorkers {
		fmt.Println("Noop HabitatWorkers")
	}
	if opt.IndustrialRobots {
		fmt.Println("Noop IndustrialRobots")
	}
	buildings := s.Data().Buildings
	for i := 0; i < len(buildings); i++ {
		if buildings[i].ConsumerProducer != nil && buildings[i].ConsumerProducer.Type == TYPE_HABITAT {
			if opt.HabitatStorage {
				buildings[i] = maxBuildingStorage(buildings[i])
			}
		}

		if opt.FactoryStorage && buildings[i].ConsumerProducer != nil && buildings[i].ConsumerProducer.Type == TYPE_FACTORY {
			buildings[i] = maxBuildingStorage(buildings[i])
		}
	}
	s.Changed = true
}
