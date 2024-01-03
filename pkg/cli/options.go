package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/heathcliff26/infraspace-savegame-editor/pkg/save"
)

var (
	version bool

	noBackup bool

	Options CLIOptions
)

type CLIOptions struct {
	Path               string
	UnlockResearch     bool
	UnlockAllResearch  bool
	Show               bool
	StarterWorkerCount IntFlag
	Resources          IntMapFlag
	Backup             bool
	BuildingOptions    save.EditBuildingsOptions
}

func init() {
	Options.Resources = make(IntMapFlag)

	flag.BoolVar(&version, "version", false, "Print version information and exit")
	flag.StringVar(&Options.Path, "p", "", "Requiered: Path to the savegame")
	flag.BoolVar(&Options.Show, "s", false, "Show the current values of the safe")

	flag.BoolVar(&Options.UnlockResearch, "research", false, "Unlock research")
	flag.BoolVar(&Options.UnlockAllResearch, "researchall", false, "Unlock all research")

	flag.Var(&Options.StarterWorkerCount, "setWorkers", "Increase the starter workers to the given value")
	flag.Var(Options.Resources, "resource", "Set the resource to the given value, can be used multiple times")

	flag.BoolVar(&Options.BuildingOptions.HabitatStorage, "maxHabitatStorage", false, "Set all resources in the habitat to 1000")
	flag.BoolVar(&Options.BuildingOptions.HabitatWorkers, "maxHabitatWorkers", false, "Fill all habitats with workers")
	flag.BoolVar(&Options.BuildingOptions.IndustrialRobots, "industrialRobots", false, "Fill all Industrial Robot factorys with 1 mio robots and resources")
	flag.BoolVar(&Options.BuildingOptions.FactoryStorage, "factoryStorage", false, "Fill all storage in factory buildings to 100")

	flag.BoolVar(&noBackup, "nobackup", false, "Do not create a backup of the save")
}

func getGitRevision(settings []debug.BuildSetting) string {
	for _, v := range settings {
		if v.Key == "vcs.revision" {
			return v.Value
		}
	}
	return "(devel)"
}

func ParseCLI() *CLIOptions {
	flag.Parse()

	if version {
		if buildinfo, ok := debug.ReadBuildInfo(); ok {
			fmt.Println("Version: " + getGitRevision(buildinfo.Settings))
			fmt.Println(buildinfo.GoVersion)
		} else {
			fmt.Println("Failed to read BuildInfo")
		}
		fmt.Println(save.GAME_VERSION)

		os.Exit(0)
	}

	exitWithError := false
	if Options.Path == "" {
		printErrorMsg("Missing or empty path")
		exitWithError = true
	}
	if Options.UnlockResearch && Options.UnlockAllResearch {
		printErrorMsg("You can't use research and researchall at the same time")
		exitWithError = true
	}
	if Options.BuildingOptions.IndustrialRobots && Options.BuildingOptions.FactoryStorage {
		printErrorMsg("You can't use industrialRobots and factoryStorage at the same time")
		exitWithError = true
	}

	if exitWithError {
		fmt.Fprint(os.Stderr, "Usage:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	Options.Backup = !noBackup
	return &Options
}

// Print message to stderr
func printErrorMsg(msg string) {
	fmt.Fprint(os.Stderr, "Error: "+msg+"\n")
}
