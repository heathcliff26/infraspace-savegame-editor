package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	pathFlag           stringFlag
	unlockResearch     bool
	starterWorkerCount intFlag
)

type stringFlag struct {
	set   bool
	value string
}

type intFlag struct {
	set   bool
	value int
}

func (flag *stringFlag) Set(arg string) error {
	flag.value = arg
	flag.set = true
	return nil
}

func (flag *stringFlag) String() string {
	return flag.value
}

func (flag *intFlag) Set(arg string) error {
	i, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("Expected a number")
	}
	flag.value = i
	flag.set = true
	return nil
}

func (flag *intFlag) String() string {
	return string(flag.value)
}

func init() {
	flag.Var(&pathFlag, "p", "Requiered: Path to the savegame")
	flag.BoolVar(&unlockResearch, "research", false, "Unlock all research")
	flag.Var(&starterWorkerCount, "setWorkers", "Increase the starter workers to the given value")
}

func main() {
	flag.Parse()

	if !pathFlag.set {
		fmt.Println("Error: Path is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	save, err := LoadSavegame(pathFlag.value)
	if err != nil {
		fmt.Printf("Could not load savegame %s: %v\n", pathFlag.value, err)
		os.Exit(1)
	}

	changed := false

	if unlockResearch {
		unlockAllResearch(save)
		fmt.Println("Unlocked all Research")
		changed = true
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

	if !changed {
		fmt.Println("There was nothing to change")
		os.Exit(1)
	}

	save.setPath("edited.sav")
	err = save.Save()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
