package save

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
)

type SaveData struct {
	MapSeed                    int64               `json:"mapSeed"`
	MapGenVersion              int64               `json:"mapGenVersion"`
	NextID                     int64               `json:"nextID"`
	SimulationFrame            int64               `json:"simulationFrame"`
	TotalGameTime              jsontext.Value      `json:"totalGameTime"`
	TotalPlayTime              jsontext.Value      `json:"totalPlayTime"`
	SaveFixGracePeriodActive   bool                `json:"saveFixGracePeriodActive"`
	WorldSettings              jsontext.Value      `json:"worldSettings"`
	Buildings                  []Building          `json:"buildings"`
	BuildingConnectors         jsontext.Value      `json:"buildingConnectors"`
	BuildingGroups             jsontext.Value      `json:"buildingGroups"`
	NetEdges                   jsontext.Value      `json:"netEdges"`
	NetNodes                   jsontext.Value      `json:"netNodes"`
	Cars                       jsontext.Value      `json:"cars"`
	Market                     Market              `json:"market"`
	Resources                  map[string]int64    `json:"resources"`
	GoalManager                jsontext.Value      `json:"goalManager"`
	ResearchManager            ResearchManager     `json:"researchManager"`
	UpgradeManager             jsontext.Value      `json:"upgradeManager,omitempty"`
	PopulationManager          jsontext.Value      `json:"populationManager"`
	Statistics                 jsontext.Value      `json:"statistics"`
	Camera                     jsontext.Value      `json:"camera"`
	DistrictManager            jsontext.Value      `json:"districtManager"`
	TrainLineManager           jsontext.Value      `json:"trainLineManager"`
	Trains                     jsontext.Value      `json:"trains"`
	CarCarriers                jsontext.Value      `json:"carCarriers"`
	Spaceship                  Spaceship           `json:"spaceship"`
	PipeComponentManager       jsontext.Value      `json:"pipeComponentManager"`
	ScriptMods                 jsontext.Value      `json:"scriptMods"`
	NewWorldPersistent         NewWorldPersistent  `json:"newWorldPersistent"`
	EnvironmentObjects         []EnvironmentObject `json:"environmentObjects"` // Very Big Object ca. 8k lines
	TerraformingProgressString jsontext.Value      `json:"terraformingProgressString"`
	AchievementsManager        AchievementsManager `json:"achievementsManager"`
	TrailerModule              jsontext.Value      `json:"trailerModule"`
	StoryMessagesModule        jsontext.Value      `json:"storyMessagesModule"`
	Roads                      jsontext.Value      `json:"roads"`
	Intersections              jsontext.Value      `json:"intersections"`
}

type Building struct {
	BuildingName            string            `json:"buildingName"`
	CustomName              jsontext.Value    `json:"customName"`
	Road                    int64             `json:"road"`
	Pipes                   jsontext.Value    `json:"pipes"`
	ID                      int64             `json:"ID"`
	Position                jsontext.Value    `json:"position"`
	Rotation                jsontext.Value    `json:"rotation"`
	ConsumerProducer        *ConsumerProducer `json:"consumerProducer"`
	MissingResourceDuration jsontext.Value    `json:"missingResourceDuration"`
	Upgrades                jsontext.Value    `json:"upgrades"`
	IntegratedNetEdges      jsontext.Value    `json:"integratedNetEdges"`
	TextModule              jsontext.Value    `json:"textModule"`
	StationModule           jsontext.Value    `json:"stationModule"`
}

type ConsumerProducer struct {
	ProductionLogic       interface{}    `json:"productionLogic"`
	IncomingStorage       []int64        `json:"incomingStorage"`
	OutgoingStorage       []int64        `json:"outgoingStorage"`
	RequestStatusDirty    bool           `json:"requestStatusDirty"`
	LastStepPowerProduced jsontext.Value `json:"lastStepPowerProduced"`
	LastStepPowerNeeded   jsontext.Value `json:"lastStepPowerNeeded"`

	Type BuildingType `json:"-"`
}

type BuildingType int

type FactoryProductionLogic struct {
	Type                 string         `json:"$type"`
	ProductionDefinition jsontext.Value `json:"productionDefinition"`
	LogicOverride        jsontext.Value `json:"logicOverride"`
	TerraformRadius      jsontext.Value `json:"terraformRadius"`
	TerraformType        jsontext.Value `json:"terraformType"`
	ProductionTimeStep   int64          `json:"productionTimeStep"`
}

type HabitatProductionLogic struct {
	Type                    string                    `json:"$type"`
	Storage                 map[string]jsontext.Value `json:"storage"`
	MaxInhabitants          int64                     `json:"maxInhabitants"`
	HabitatLevel            int64                     `json:"habitatLevel"`
	Upgrade                 jsontext.Value            `json:"upgrade"`
	Downgrade               jsontext.Value            `json:"downgrade"`
	PowerNeededForTenPeople jsontext.Value            `json:"powerNeededForTenPeople"`
	TargetInhabitants       jsontext.Value            `json:"targetInhabitants"`
	UpgradeCountdown        jsontext.Value            `json:"upgradeCountdown"`
	DowngradeCountdown      jsontext.Value            `json:"downgradeCountdown"`
	Workers                 []Worker                  `json:"workers"`
}

type Market struct {
	StarterWorkers     []Worker       `json:"starterWorkers"`
	ResourcePriorities jsontext.Value `json:"resourcePriorities"`
}

type Worker struct {
	Home int64 `json:"_home"`
	ID   int64
}

type ResearchManager struct {
	ResearchProgress map[string]int64 `json:"researchProgress"`
	CurrentResearch  jsontext.Value   `json:"currentResearch"`
	ResearchQueue    []string         `json:"researchQueue"`
}

type Spaceship struct {
	CurrentlyRepairedPartName jsontext.Value `json:"currentlyRepairedPartName"`
	Parts                     []struct {
		Type           string         `json:"$type,omitempty"`
		TargetPosition jsontext.Value `json:"targetPosition,omitempty"`
		Timer          jsontext.Value `json:"timer,omitempty"`
		Name           string         `json:"name"`
		CurrentSteps   int64          `json:"currentSteps"`
	} `json:"parts"`
	CurrentlyRepairedPartNameKBackingField jsontext.Value `json:"<currentlyRepairedPartName>k__BackingField"`
}

type NewWorldPersistent struct {
	HeightData string         `json:"heightData"`
	AlphaData  []string       `json:"alphaData"`
	DetailData []string       `json:"detailData"`
	BiomesData jsontext.Value `json:"biomesData"`
}

type EnvironmentObject struct {
	ID                  int64
	ObjectName          string         `json:"objectName"`
	Health              jsontext.Value `json:"health"`
	TransformCompressed string         `json:"transformCompressed"`
}

type AchievementsManager struct {
	UnlockabilityStatus           UnlockabilityStatus `json:"unlockabilityStatus"`
	SerializedAchievementTrackers jsontext.Value      `json:"serializedAchievementTrackers"`
}

type UnlockabilityStatus struct {
	DisabledDueToMods                             bool `json:"disabledDueToMods"`
	DisabledDueToCreativeSettings                 bool `json:"disabledDueToCreativeSettings"`
	DisabledDueToSettingsModification             bool `json:"disabledDueToSettingsModification"`
	DisabledDueToCheats                           bool `json:"disabledDueToCheats"`
	DisabledDueToModsBackingField                 bool `json:"<disabledDueToMods>k__BackingField"`
	DisabledDueToCreativeSettingsBackingField     bool `json:"<disabledDueToCreativeSettings>k__BackingField"`
	DisabledDueToSettingsModificationBackingField bool `json:"<disabledDueToSettingsModification>k__BackingField"`
	DisabledDueToCheatsBackingField               bool `json:"<disabledDueToCheats>k__BackingField"`
}

func (c *ConsumerProducer) UnmarshalJSON(data []byte) error {
	var v map[string]jsontext.Value
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	err = json.Unmarshal(v["incomingStorage"], &c.IncomingStorage)
	if err != nil {
		return err
	}
	err = json.Unmarshal(v["outgoingStorage"], &c.OutgoingStorage)
	if err != nil {
		return err
	}
	err = json.Unmarshal(v["requestStatusDirty"], &c.RequestStatusDirty)
	if err != nil {
		return err
	}
	err = json.Unmarshal(v["lastStepPowerProduced"], &c.LastStepPowerProduced)
	if err != nil {
		return err
	}
	err = json.Unmarshal(v["lastStepPowerNeeded"], &c.LastStepPowerNeeded)
	if err != nil {
		return err
	}

	var tmp struct {
		Type string `json:"$type"`
	}
	err = json.Unmarshal(v["productionLogic"], &tmp)
	if err != nil {
		return err
	}
	switch tmp.Type {
	case "FactoryProductionLogic, old":
		c.Type = TYPE_FACTORY
		var fProd FactoryProductionLogic
		err = json.Unmarshal(v["productionLogic"], &fProd)
		if err != nil {
			return err
		}
		c.ProductionLogic = fProd
	case "Habitat, old":
		c.Type = TYPE_HABITAT
		var hProd HabitatProductionLogic
		err = json.Unmarshal(v["productionLogic"], &hProd)
		if err != nil {
			return err
		}
		c.ProductionLogic = hProd
	case "StorehouseProductionLogic, old":
		c.Type = TYPE_STOREHOUSE
	case "ResearchProductionLogic, old":
		c.Type = TYPE_RESEARCH
	default:
		c.Type = TYPE_UNKNOWN
	}
	if c.Type > 1 {
		var raw jsontext.Value
		err = json.Unmarshal(v["productionLogic"], &raw)
		if err != nil {
			return err
		}
		c.ProductionLogic = raw
	}

	return nil
}
