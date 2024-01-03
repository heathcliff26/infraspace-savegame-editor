package main

import (
	"fmt"
	"os"

	"github.com/heathcliff26/infraspace-savegame-editor/pkg/save"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Receveived wrong number of arguments, expected 1 but got %d", len(os.Args))
		os.Exit(1)
	}
	s, err := save.LoadSavegame(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load savegame: %v\n", err)
		os.Exit(1)
	}

	// This is about 34 MB
	s.Data().NewWorldPersistent.HeightData = ""

	// There can be entries with about 1MB apiece in these arrays
	s.Data().NewWorldPersistent.AlphaData = trimStrings(s.Data().NewWorldPersistent.AlphaData)
	s.Data().NewWorldPersistent.DetailData = trimStrings(s.Data().NewWorldPersistent.DetailData)

	bPath, err := s.Backup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to backup savegame: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Created backup of save file at " + bPath)

	err = s.Save()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load savegame: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully trimmed savefile")
}

func trimStrings(list []string) []string {
	d := 0
	for i := range list {
		if len(list[i-d]) > 50 {
			if len(list)-1 == i-d {
				list = list[:i-d]
			} else {
				list = append(list[:i-d], list[i+1-d:]...)
			}
			d++
		}
	}
	return list
}
