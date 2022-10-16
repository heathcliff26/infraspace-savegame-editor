package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const GAME_VERSION = "InfraSpace Alpha 8.1.184"

var (
	version            bool
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
	flag.BoolVar(&version, "version", false, "Print version information and exit")
	flag.Var(&pathFlag, "p", "Requiered: Path to the savegame")
	flag.BoolVar(&show, "s", false, "Show the current values of the safe")
	flag.BoolVar(&research, "research", false, "Unlock research")
	flag.BoolVar(&researchall, "researchall", false, "Unlock all research")
	flag.Var(&starterWorkerCount, "setWorkers", "Increase the starter workers to the given value")
	flag.Var(resourceFlags, "resource", "Set the resource to the given value, can be used multiple times")
}

func parseFlags() {
	flag.Parse()

	if version {
		fmt.Printf("This editor was made with %s\n", GAME_VERSION)
		os.Exit(0)
	}
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
