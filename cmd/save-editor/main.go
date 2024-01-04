package main

import (
	"github.com/heathcliff26/infraspace-savegame-editor/pkg/gui"
)

func main() {
	app := gui.New()
	app.Run()
	// opt := cli.ParseCLI()

	// s, err := save.LoadSavegame(opt.Path)
	// if err != nil {
	// 	exitWithError(fmt.Sprintf("Could not load savegame %s: %s", opt.Path, err))
	// }

	// if opt.Show {
	// 	s.Print()
	// }

	// if opt.UnlockAllResearch {
	// 	s.UnlockAllResearch()
	// 	fmt.Println("Unlocked all Research")
	// } else if opt.UnlockResearch {
	// 	fmt.Println("Noop UnlockResearch")
	// 	// err = unlockResearch(s)
	// 	// if err != nil {
	// 	// 	exitWithError(fmt.Sprintf("Could not unlock research: %s", err))
	// 	// } else {
	// 	// 	fmt.Printf("Unlocked Research\n")
	// 	// 	changed = true
	// 	// }
	// }
	// if opt.StarterWorkerCount.IsSet() {
	// 	numAdded := s.AddStarterWorkers(opt.StarterWorkerCount.Value)
	// 	if numAdded > 0 {
	// 		fmt.Printf("Increased starter workers by %d\n", numAdded)
	// 	} else {
	// 		fmt.Println("Could not increase the starter workers, there where already more workers present")
	// 	}
	// }
	// for key, value := range opt.Resources {
	// 	err = s.SetResource(key, value)
	// 	if err != nil {
	// 		exitWithError(fmt.Sprintf("Could not set resource: %v", err))
	// 	} else {
	// 		fmt.Printf("Set %s to %d\n", key, value)
	// 	}
	// }
	// bOpt := opt.BuildingOptions
	// if bOpt.HabitatStorage || bOpt.HabitatWorkers || bOpt.IndustrialRobots || bOpt.FactoryStorage {
	// 	s.EditBuildings(opt.BuildingOptions)
	// }

	// if !s.Changed && !opt.Show {
	// 	fmt.Printf("There was nothing to change\n")
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// } else if s.Changed {
	// 	if opt.Backup {
	// 		dst, err := s.Backup()
	// 		if err != nil {
	// 			exitWithError(fmt.Sprintf("Could not backup save: %s", err))
	// 		}
	// 		fmt.Printf("Created backup of save at %s\n", dst)
	// 	}
	// 	err = s.Save()
	// 	if err != nil {
	// 		exitWithError(fmt.Sprintf("Could not save changes: %s", err))
	// 	} else {
	// 		fmt.Printf("Changes have been written to %s\n", s.Path())
	// 	}
	// }
}
