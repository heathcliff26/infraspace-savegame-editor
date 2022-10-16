package main

import (
	"flag"
	"fmt"
	"os"
)

func exitError(msg string) {
	fmt.Println(msg)
	fmt.Println("No changes have been written to the file")
	os.Exit(1)
}

func main() {
	parseFlags()

	save, err := LoadSavegame(pathFlag.value)
	if err != nil {
		exitError(fmt.Sprintf("Could not load savegame %s: %s", pathFlag.value, err))
	}

	if show {
		printSaveInfo(save)
	}

	changed := false

	if researchall {
		unlockAllResearch(save)
		fmt.Printf("Unlocked all Research\n")
		changed = true
	} else if research {
		err = unlockResearch(save)
		if err != nil {
			exitError(fmt.Sprintf("Error: Could not unlock research: %s", err))
		} else {
			fmt.Printf("Unlocked Research\n")
			changed = true
		}
	}
	if starterWorkerCount.set {
		count := increaseStarterWorkers(save, starterWorkerCount.value)
		if count == starterWorkerCount.value {
			fmt.Printf("Increased starter workers to %d\n", count)
			changed = true
		} else {
			fmt.Printf("Could not increase the starter workers, there where already %d workers\n", count)
		}
	}
	for name, resourceFlag := range resourceFlags {
		if resourceFlag.set {
			err = setResource(save, name, resourceFlag.value)
			if err != nil {
				exitError(fmt.Sprintf("Could not set resource %s: %s", name, err))
			} else {
				fmt.Printf("Set %s to %d\n", name, resourceFlag.value)
				changed = true
			}
		}
	}
	if habitatStorage {
		maxHabitatStorage(save)
		fmt.Printf("Set Habitat Storage to 1000\n")
		changed = true
	}

	if !changed && !show {
		fmt.Printf("There was nothing to change\n")
		flag.PrintDefaults()
		os.Exit(1)
	} else if changed {
		save.setPath("edited.sav")
		err = save.Save()
		if err != nil {
			exitError(fmt.Sprintf("Error: %s", err))
		} else {
			fmt.Printf("Changes have been writtent to %s\n", save.getPath())
		}
	}
}
