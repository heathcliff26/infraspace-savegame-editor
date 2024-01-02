package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var researchIndex = []string{
	"farming",
	"electronics",
	"steelMaking",
	"oneWayRoads",
	"science2",
	"homeAppliances",
	"solarPanels",
	"improvedWindTurbineBlades",
	"concreteRoads",
	"aluminiumMining",
	"powerSubstation",
	"motors",
	"fourLaneRoads",
	"parks",
	"computers",
	"largeMines",
	"highways",
	"trains",
	"foodProcessing",
	"fertilizer",
	"nanotubes",
	"selfCleaningSolarPanels",
	"sixLaneRoads",
	"stadiums",
	"neuralProcessors",
	"science3",
	"uraniumMining",
	"superhighways",
	"homeRobots",
	"industrialRobots",
	"particleFiltering",
	"nuclearPower",
	"schools",
	"iridiumMining",
	"recycling",
	"holoDisplays",
	"highTechWorkshop",
	"aiControlUnits",
	"iridiumPropellant",
	"fastNeutronReactor",
	"science4",
	"adamantineMining",
}

var maxResearchProgress = map[string]int{
	"farming":                   5,
	"electronics":               15,
	"steelMaking":               10,
	"oneWayRoads":               15,
	"science2":                  40,
	"homeAppliances":            20,
	"solarPanels":               25,
	"improvedWindTurbineBlades": 15,
	"concreteRoads":             20,
	"aluminiumMining":           40,
	"powerSubstation":           50,
	"motors":                    40,
	"fourLaneRoads":             40,
	"parks":                     80,
	"computers":                 60,
	"largeMines":                60,
	"highways":                  60,
	"trains":                    100,
	"foodProcessing":            60,
	"fertilizer":                80,
	"nanotubes":                 30,
	"selfCleaningSolarPanels":   50,
	"sixLaneRoads":              50,
	"stadiums":                  100,
	"neuralProcessors":          70,
	"science3":                  120,
	"uraniumMining":             40,
	"superhighways":             120,
	"homeRobots":                90,
	"industrialRobots":          90,
	"particleFiltering":         60,
	"nuclearPower":              80,
	"schools":                   200,
	"iridiumMining":             80,
	"recycling":                 120,
	"holoDisplays":              150,
	"highTechWorkshop":          180,
	"aiControlUnits":            220,
	"iridiumPropellant":         180,
	"fastNeutronReactor":        200,
	"science4":                  350,
	"adamantineMining":          350,
}

var habitatResourceNames = []string{
	"oxygen",
	"survivalFood",
	"homeAppliance",
	"computer",
	"parkPoints",
	"goodFood",
	"homeRobot",
	"culturePoints",
}

const RESOURCE_FACTOR = 100

type worker struct {
	_home int
	ID    int
}

// unlock all research
func unlockAllResearch(save *savegame) {
	researchProgress := save.getResearchProgress()
	for key := range researchProgress {
		researchProgress[key] = float64(maxResearchProgress[key])
	}
}

// interactively unlock specific research
func unlockResearch(save *savegame) error {
	researchProgress := save.getResearchProgress()
	lockedResearch := make([]string, 0, len(maxResearchProgress))
	i := 0
	fmt.Printf("Research to unlock:\n")
	for x := 0; x < len(researchIndex); x++ {
		key := researchIndex[x]
		value := researchProgress[key]
		if value.(float64) < float64(maxResearchProgress[key]) {
			lockedResearch = append(lockedResearch, key)
			fmt.Printf("\t%d.\t%s\n", i, key)
			i++
		}
	}
	fmt.Printf("\nEnter the numbers of the items to unlock and press enter\n> ")

	r := bufio.NewReader(os.Stdin)

	input, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Printf("Input: %s\nUnlocking the following research:\n", input)
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, "\r")
	indexes := strings.Split(input, " ")
	for _, index := range indexes {
		i, err = strconv.Atoi(index)
		if err != nil {
			return err
		}
		if i < 0 || i >= len(lockedResearch) {
			return fmt.Errorf("Invalid input, %d is not a valid choice", i)
		}
		key := lockedResearch[i]
		researchProgress[key] = float64(maxResearchProgress[key])
		fmt.Printf("\t%s\n", key)
	}

	return nil
}

// increase the starter workers ot the given count, return resulting number of starting workers
func increaseStarterWorkers(save *savegame, count int) int {
	starterWorkers := save.getStarterWorkers()
	nextID := save.getNextID()
	var newWorker worker
	for len(starterWorkers) < count {
		newWorker = worker{
			_home: 0,
			ID:    nextID,
		}
		nextID++
		starterWorkers = append(starterWorkers, newWorker)
	}
	save.setNextID(nextID)
	return len(starterWorkers)
}

// Print all information about the editable data in the savegame
func printSaveInfo(save *savegame) {
	metadata := strings.Split(save.getPrefix(), "\n")
	fmt.Printf("Printing savegame Information...\n\nMetadata:\n\tGame Version: %s\n\tCreated: %s\n\n", metadata[1], metadata[2])

	resources := save.getResources()
	fmt.Printf("Resources:\n")
	for key, value := range resources {
		resource := int(value.(float64) / RESOURCE_FACTOR)
		fmt.Printf("\t%s: %d\n", key, resource)
	}
	fmt.Printf("\n")

	starterWorkers := save.getStarterWorkers()
	fmt.Printf("starter workers: %d\n\n", len(starterWorkers))

	buildings := save.getBuildings()
	fmt.Printf("building count: %d\n\n", len(buildings))

	researchProgress := save.getResearchProgress()
	fmt.Printf("Unlocked Research:\n")
	researchProgressEmpty := true
	for i := 0; i < len(researchIndex); i++ {
		key := researchIndex[i]
		value := researchProgress[key]
		if value.(float64) == float64(maxResearchProgress[key]) {
			fmt.Printf("\t%s\n", key)
			researchProgressEmpty = false
		}
	}
	if researchProgressEmpty {
		fmt.Printf("\tnone\n\n")
	} else {
		fmt.Printf("\n")
	}
}

func setResource(save *savegame, resource string, value int) error {
	resources := save.getResources()
	if _, ok := resources[resource]; !ok {
		return fmt.Errorf("Unkown resource %s", resource)
	}
	resources[resource] = float64(value * RESOURCE_FACTOR)
	return nil
}

func getBuildingProductionLogic(building map[string]interface{}) (map[string]interface{}, bool) {
	return GetJSONObject(building, "consumerProducer", "productionLogic")
}

func checkIfHabitat(building map[string]interface{}) bool {
	buildingName := building["buildingName"].(string)
	return strings.HasPrefix(buildingName, "habitatLevel")
}

func checkIfIndustrialRobotFactory(building map[string]interface{}) bool {
	buildingName := building["buildingName"].(string)
	return buildingName == "industrialRobotFactory"
}

func maxHabitatStorage(building map[string]interface{}) error {
	if checkIfHabitat(building) {
		return nil
	}
	storage, ok := GetJSONObject(building, "consumerProducer", "productionLogic", "storage")
	if !ok {
		return fmt.Errorf("Could not fill Habitat Storage")
	}
	for _, resourceName := range habitatResourceNames {
		storage[resourceName] = 10
	}
	return nil
}

func maxHabitatWorkers(building map[string]interface{}, nextID int) (int, error) {
	if checkIfHabitat(building) {
		return nextID, nil
	}
	productionLogic, ok := getBuildingProductionLogic(building)
	if !ok {
		return nextID, fmt.Errorf("Could not max habitat workers")
	}
	maxInhabitants := int(productionLogic["maxInhabitants"].(float64))
	homeID := int(building["ID"].(float64))
	workers, ok := GetJSONArray(productionLogic, "workers")
	if !ok {
		return nextID, fmt.Errorf("Could not max habitat workers")
	}
	var newWorker worker
	for len(workers) < maxInhabitants {
		newWorker = worker{
			_home: homeID,
			ID:    nextID,
		}
		nextID++
		workers = append(workers, newWorker)
	}
	return nextID, nil
}

func setFactoryStorage(factory map[string]interface{}, targetAmount int) error {
	incomingStorage, ok := GetJSONArray(factory, "consumerProducer", "incomingStorage")
	if !ok {
		return fmt.Errorf("Could not find storage")
	}
	// skip last entry, as it is not needed
	for i := 0; len(incomingStorage)-1 > i; i++ {
		incomingStorage[i] = targetAmount
	}
	outgoingStorage, ok := GetJSONArray(factory, "consumerProducer", "outgoingStorage")
	if !ok {
		return fmt.Errorf("Could not find storage")
	}
	for i := 0; len(outgoingStorage) > i; i++ {
		outgoingStorage[i] = targetAmount
	}
	return nil
}

func maxFactoryStorage(building map[string]interface{}) error {
	productionLogic, ok := getBuildingProductionLogic(building)
	if !ok {
		return nil
	}
	buildingType := productionLogic["$type"].(string)
	if !strings.HasPrefix(buildingType, "FactoryProductionLogic") {
		return nil
	}
	if checkIfIndustrialRobotFactory(building) {
		return setFactoryStorage(building, 1000000)
	} else {
		return setFactoryStorage(building, 100)
	}
}

func maxIndustrialRobots(building map[string]interface{}) error {
	if !checkIfIndustrialRobotFactory(building) {
		return nil
	}
	return setFactoryStorage(building, 1000000)
}

func editBuildings(save *savegame, args map[string]bool) error {
	buildings := save.getBuildings()
	nextID := save.getNextID()
	oldNextID := nextID
	habitatWorkers := args["habitatWorkers"]
	habitatStorage := args["habitatStorage"]
	industrialRobots := args["industrialRobots"]
	factoryStorage := args["factoryStorage"]
	var err error
	for i := 0; i < len(buildings); i++ {
		building := buildings[i].(map[string]interface{})
		if habitatWorkers {
			nextID, err = maxHabitatWorkers(building, nextID)
			if err != nil {
				return err
			}
		}
		if habitatStorage {
			err = maxHabitatStorage(building)
			if err != nil {
				return err
			}
		}
		if factoryStorage {
			err = maxFactoryStorage(building)
			if err != nil {
				return err
			}
		}
		if industrialRobots {
			err = maxIndustrialRobots(building)
			if err != nil {
				return err
			}
		}
	}
	if nextID != oldNextID {
		save.setNextID(nextID)
	}
	return nil
}
