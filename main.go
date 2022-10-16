package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	pathFlag           stringFlag
	research           bool
	researchall        bool
	show               bool
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
	flag.BoolVar(&show, "s", false, "Show the current values of the safe")
	flag.BoolVar(&research, "research", false, "Unlock research")
	flag.BoolVar(&researchall, "researchall", false, "Unlock all research")
	flag.Var(&starterWorkerCount, "setWorkers", "Increase the starter workers to the given value")
}

func parseFlags() {
	flag.Parse()

	exitWithError := false
	if !pathFlag.set {
		fmt.Println("Error: Path is required")
		exitWithError = true
	}
	if research && researchall {
		fmt.Println("Error: You can't use research and researchall at the same time")
		exitWithError = true
	}

	if exitWithError {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	parseFlags()

	save, err := LoadSavegame(pathFlag.value)
	if err != nil {
		fmt.Printf("Could not load savegame %s: %s\n", pathFlag.value, err)
		os.Exit(1)
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
			fmt.Printf("Error: Could not unlock research: %s\n", err)
			os.Exit(1)
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

	if !changed && !show {
		fmt.Printf("There was nothing to change\n")
		flag.PrintDefaults()
		os.Exit(1)
	} else if changed {
		save.setPath("edited.sav")
		err = save.Save()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Changes have been writtent to %s\n", save.getPath())
		}
	}
}
