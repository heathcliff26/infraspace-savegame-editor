package save

const (
	GAME_VERSION         = "1.33.423"
	RESOURCE_FACTOR      = 100
	BUILDING_MAX_STORAGE = 100
)

const (
	TYPE_FACTORY    BuildingType = iota
	TYPE_HABITAT    BuildingType = iota
	TYPE_STOREHOUSE BuildingType = iota
	TYPE_RESEARCH   BuildingType = iota
	TYPE_UNKNOWN    BuildingType = iota
)

var maxResearchProgress = map[string]uint{
	"farming":                   10,
	"electronics":               40,
	"steelMaking":               20,
	"oneWayRoads":               15,
	"pipes":                     40,
	"distribution":              50,
	"aluminiumMining":           40,
	"solarPanels":               65,
	"improvedWindTurbineBlades": 25,
	"concreteRoads":             40,
	"fence":                     50,
	"methane":                   50,
	"tanks":                     50,
	"computers":                 75,
	"science2":                  50,
	"powerSubstation":           100,
	"spaceshipConstruction":     150,
	"fourLaneRoads":             65,
	"cargoGondolas":             100,
	"pixelBuilding":             50,
	"methanePowerPlants":        100,
	"parks":                     100,
	"groundWaterExtraction":     150,
	"motors":                    75,
	"highways":                  100,
	"trains":                    200,
	"foodProcessing":            75,
	"particleFiltering":         125,
	"nanotubes":                 100,
	"selfCleaningSolarPanels":   200,
	"largeMines":                150,
	"sixLaneRoads":              100,
	"statue":                    125,
	"methaneFermentation":       250,
	"science3":                  200,
	"neuralProcessors":          100,
	"culture":                   150,
	"uraniumMining":             75,
	"superhighways":             250,
	"sightSeeing1":              150,
	"homeRobots":                150,
	"industrialRobots":          100,
	"nuclearPower":              150,
	"recycling":                 125,
	"highSpeedRail":             100,
	"sightSeeing2":              150,
	"schools":                   200,
	"iridiumMining":             100,
	"fertilizer":                100,
	"heatExchanger":             125,
	"holoDisplays":              150,
	"aiControlUnits":            200,
	"highTechWorkshop":          250,
	"science4":                  350,
	"superconductor":            100,
	"lightningRail":             150,
	"adamantineMining":          250,
	"fastNeutronReactor":        200,
	"terraforming":              150,
	"droneFertilization":        150,
	"groundAcidityRegulation":   250,
	"ammoniaExtraction":         350,
	"magneticFieldGeneration":   500,
}
