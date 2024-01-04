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

func (s *Savegame) Path() string {
	return s.path
}

func (s *Savegame) Prefix() string {
	return s.prefix
}

func (s *Savegame) Data() *SaveData {
	return s.data
}

// Get a resource by name
func (s *Savegame) GetResource(key string) (int, bool) {
	value, ok := s.Data().Resources[key]
	return int(value / RESOURCE_FACTOR), ok
}

// Get unlocked research
func (s *Savegame) GetUnlockedResearch() []string {
	res := make([]string, 0, len(maxResearchProgress))
	for k, v := range maxResearchProgress {
		if s.Data().ResearchManager.ResearchProgress[k] == v {
			res = append(res, k)
		}
	}
	return res
}

// Get the current number of starter workers
func (s *Savegame) GetStarterWorkerCount() int {
	return len(s.Data().Market.StarterWorkers)
}

// Unlock research by name
func (s *Savegame) UnlockResearch(name string) {
	if s.Data().ResearchManager.ResearchProgress[name] != maxResearchProgress[name] {
		s.Data().ResearchManager.ResearchProgress[name] = maxResearchProgress[name]
		s.Changed = true
	}
}

// Lock Research by name, if currently unlocked
func (s *Savegame) LockResearch(name string) {
	if s.Data().ResearchManager.ResearchProgress[name] == maxResearchProgress[name] {
		s.Data().ResearchManager.ResearchProgress[name] = 0
		s.Changed = true
	}
}

// Unlocks all research
func (s *Savegame) UnlockAllResearch() {
	for name := range s.Data().ResearchManager.ResearchProgress {
		s.UnlockResearch(name)
	}
}

// Increase the starter workers to the given count, return number of added workers
func (s *Savegame) AddStarterWorkers(count int) int64 {
	oldNextID := s.Data().NextID
	for s.GetStarterWorkerCount() < count {
		newWorker := Worker{
			Home: 0,
			ID:   s.Data().NextID,
		}
		s.Data().NextID++
		s.Data().Market.StarterWorkers = append(s.Data().Market.StarterWorkers, newWorker)
	}

	diff := (s.Data().NextID - oldNextID)
	if diff > 0 {
		s.Changed = true
	}
	return diff
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
