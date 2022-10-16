package main

import (
	"fmt"
	"strings"
)

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

type worker struct {
	_home int
	ID    int
}

// unlock all research
func unlockAllResearch(save *savegame) {
	researchProgress := save.getResearchProgress()
	for key, _ := range researchProgress {
		researchProgress[key] = float64(maxResearchProgress[key])
	}
}

// increase the starter workers ot the given count, return resulting number of starting workers
func increaseStarterWorkers(save *savegame, count int) int {
	starterWorkers := save.getStarterWorkers()
	var nextID int = int(save.Data()["nextID"].(float64))
	var newWorker worker
	for len(starterWorkers) < count {
		newWorker = worker{
			_home: 0,
			ID:    nextID,
		}
		nextID++
		starterWorkers = append(starterWorkers, newWorker)
	}
	save.Data()["market"].(map[string]interface{})["starterWorkers"] = starterWorkers
	save.Data()["nextID"] = float64(nextID)
	return len(starterWorkers)
}

// Print all information about the editable data in the savegame
func printSaveInfo(save *savegame) {
	metadata := strings.Split(save.getPrefix(), "\n")
	fmt.Printf("Printing savegame Information...\n\nMetadata:\n\tGame Version: %s\n\tCreated: %s\n\n", metadata[1], metadata[2])

	resources := save.getResources()
	fmt.Printf("Resources:\n")
	for key, value := range resources {
		resource := int(value.(float64) / 10)
		fmt.Printf("\t%s: %d\n", key, resource)
	}
	fmt.Printf("\n")

	starterWorkers := save.getStarterWorkers()
	fmt.Printf("starter workers: %d\n\n", len(starterWorkers))

	researchProgress := save.getResearchProgress()
	fmt.Printf("Unlocked Research:\n")
	researchProgressEmpty := true
	for key, value := range researchProgress {
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
