package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	pathFlag           stringFlag
	research           bool
	researchall        bool
	show               bool
	starterWorkerCount intFlag
	resourceFlags      = make(intFlagMap)
)

type stringFlag struct {
	set   bool
	value string
}

type intFlag struct {
	set   bool
	value int
}

type intFlagMap map[string]intFlag

func (f *stringFlag) Set(arg string) error {
	f.value = arg
	f.set = true
	return nil
}

func (f *stringFlag) String() string {
	return f.value
}

func (f *intFlag) Set(arg string) error {
	i, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("Expected a number")
	}
	f.value = i
	f.set = true
	return nil
}

func (f *intFlag) String() string {
	return string(f.value)
}

func (f intFlagMap) Set(arg string) error {
	keyvalue := strings.Split(arg, "=")
	if len(keyvalue) != 2 {
		return fmt.Errorf("Unexcpected input '%s' for resource, expected <resource>=<value>", arg)
	}
	key := keyvalue[0]
	if _, ok := f[key]; ok {
		return fmt.Errorf("Got multiple inputs for %s", keyvalue[0])
	}
	value := intFlag{}
	err := value.Set(keyvalue[1])
	if err != nil {
		return err
	}
	f[key] = value
	return nil
}

func (f intFlagMap) String() string {
	return fmt.Sprintf("Can't convert intFlagMap to string")
}

func init() {
	flag.Var(&pathFlag, "p", "Requiered: Path to the savegame")
	flag.BoolVar(&show, "s", false, "Show the current values of the safe")
	flag.BoolVar(&research, "research", false, "Unlock research")
	flag.BoolVar(&researchall, "researchall", false, "Unlock all research")
	flag.Var(&starterWorkerCount, "setWorkers", "Increase the starter workers to the given value")
	flag.Var(resourceFlags, "resource", "Set the resource to the given value, can be used multiple times")
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
	for name, resourceFlag := range resourceFlags {
		if resourceFlag.set {
			err = setResource(save, name, resourceFlag.value)
			if err != nil {
				fmt.Printf("Could not set resource %s: %s\n", name, err)
				os.Exit(1)
			} else {
				fmt.Printf("Set %s to %d\n", name, resourceFlag.value)
				changed = true
			}
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
