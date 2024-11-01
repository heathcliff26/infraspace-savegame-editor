package save

import (
	"encoding/json"
)

type SaveData struct {
	MapSeed                    int64               `json:"mapSeed"`
	MapGenVersion              int64               `json:"mapGenVersion"`
	NextID                     int64               `json:"nextID"`
	SimulationFrame            int64               `json:"simulationFrame"`
	TotalGameTime              json.Number         `json:"totalGameTime"`
	TotalPlayTime              json.Number         `json:"totalPlayTime"`
	SaveFixGracePeriodActive   bool                `json:"saveFixGracePeriodActive"`
	WorldSettings              json.RawMessage     `json:"worldSettings"`
	Buildings                  []Building          `json:"buildings"`
	BuildingConnectors         json.RawMessage     `json:"buildingConnectors"`
	BuildingGroups             json.RawMessage     `json:"buildingGroups"`
	NetEdges                   json.RawMessage     `json:"netEdges"`
	NetNodes                   json.RawMessage     `json:"netNodes"`
	Cars                       json.RawMessage     `json:"cars"`
	Market                     Market              `json:"market"`
	Resources                  map[string]int64    `json:"resources"`
	GoalManager                json.RawMessage     `json:"goalManager"`
	ResearchManager            ResearchManager     `json:"researchManager"`
	UpgradeManager             json.RawMessage     `json:"upgradeManager,omitempty"`
	PopulationManager          json.RawMessage     `json:"populationManager"`
	Statistics                 json.RawMessage     `json:"statistics"`
	Camera                     json.RawMessage     `json:"camera"`
	DistrictManager            json.RawMessage     `json:"districtManager"`
	TrainLineManager           json.RawMessage     `json:"trainLineManager"`
	Trains                     json.RawMessage     `json:"trains"`
	CarCarriers                json.RawMessage     `json:"carCarriers"`
	Spaceship                  Spaceship           `json:"spaceship"`
	PipeComponentManager       json.RawMessage     `json:"pipeComponentManager"`
	ScriptMods                 json.RawMessage     `json:"scriptMods"`
	NewWorldPersistent         NewWorldPersistent  `json:"newWorldPersistent"`
	EnvironmentObjects         []EnvironmentObject `json:"environmentObjects"` // Very Big Object ca. 8k lines
	TerraformingProgressString json.RawMessage     `json:"terraformingProgressString"`
	AchievementsManager        AchievementsManager `json:"achievementsManager"`
	TrailerModule              json.RawMessage     `json:"trailerModule"`
	StoryMessagesModule        json.RawMessage     `json:"storyMessagesModule"`
	Roads                      json.RawMessage     `json:"roads"`
	Intersections              json.RawMessage     `json:"intersections"`
}

type Building struct {
	BuildingName            string          `json:"buildingName"`
	CustomName              json.RawMessage `json:"customName"`
	Road                    int64           `json:"road"`
	Pipes                   json.RawMessage `json:"pipes"`
	ID                      int
	Position                json.RawMessage   `json:"position"`
	Rotation                json.Number       `json:"rotation"`
	ConsumerProducer        *ConsumerProducer `json:"consumerProducer"`
	MissingResourceDuration json.Number       `json:"missingResourceDuration"`
	Upgrades                json.RawMessage   `json:"upgrades"`
	IntegratedNetEdges      json.RawMessage   `json:"integratedNetEdges"`
	TextModule              json.RawMessage   `json:"textModule"`
	StationModule           json.RawMessage   `json:"stationModule"`
}

type ConsumerProducer struct {
	ProductionLogic       interface{} `json:"productionLogic"`
	IncomingStorage       []int64     `json:"incomingStorage"`
	OutgoingStorage       []int64     `json:"outgoingStorage"`
	RequestStatusDirty    bool        `json:"requestStatusDirty"`
	LastStepPowerProduced json.Number `json:"lastStepPowerProduced"`
	LastStepPowerNeeded   json.Number `json:"lastStepPowerNeeded"`

	Type BuildingType `json:"-"`
}

type BuildingType int

type FactoryProductionLogic struct {
	Type                 string          `json:"$type"`
	ProductionDefinition json.RawMessage `json:"productionDefinition"`
	LogicOverride        json.RawMessage `json:"logicOverride"`
	TerraformRadius      json.Number     `json:"terraformRadius"`
	TerraformType        json.RawMessage `json:"terraformType"`
	ProductionTimeStep   int64           `json:"productionTimeStep"`
}

type HabitatProductionLogic struct {
	Type                    string                 `json:"$type"`
	Storage                 map[string]json.Number `json:"storage"`
	MaxInhabitants          int64                  `json:"maxInhabitants"`
	HabitatLevel            int64                  `json:"habitatLevel"`
	Upgrade                 json.RawMessage        `json:"upgrade"`
	Downgrade               json.RawMessage        `json:"downgrade"`
	PowerNeededForTenPeople json.Number            `json:"powerNeededForTenPeople"`
	TargetInhabitants       json.Number            `json:"targetInhabitants"`
	UpgradeCountdown        json.Number            `json:"upgradeCountdown"`
	DowngradeCountdown      json.Number            `json:"downgradeCountdown"`
	Workers                 []Worker               `json:"workers"`
}

type Market struct {
	StarterWorkers     []Worker        `json:"starterWorkers"`
	ResourcePriorities json.RawMessage `json:"resourcePriorities"`
}

type Worker struct {
	Home int64 `json:"_home"`
	ID   int64
}

type ResearchManager struct {
	ResearchProgress map[string]int64 `json:"researchProgress"`
	CurrentResearch  json.RawMessage  `json:"currentResearch"`
	ResearchQueue    []string         `json:"researchQueue"`
}

type Spaceship struct {
	CurrentlyRepairedPartName json.RawMessage `json:"currentlyRepairedPartName"`
	Parts                     []struct {
		Type           string          `json:"$type,omitempty"`
		TargetPosition json.RawMessage `json:"targetPosition,omitempty"`
		Timer          json.Number     `json:"timer,omitempty"`
		Name           string          `json:"name"`
		CurrentSteps   int64           `json:"currentSteps"`
	} `json:"parts"`
	CurrentlyRepairedPartNameKBackingField json.RawMessage `json:"<currentlyRepairedPartName>k__BackingField"`
}

type NewWorldPersistent struct {
	HeightData string          `json:"heightData"`
	AlphaData  []string        `json:"alphaData"`
	DetailData []string        `json:"detailData"`
	BiomesData json.RawMessage `json:"biomesData"`
}

type EnvironmentObject struct {
	ID                  int64
	ObjectName          string      `json:"objectName"`
	Health              json.Number `json:"health"`
	TransformCompressed string      `json:"transformCompressed"`
}

type AchievementsManager struct {
	UnlockabilityStatus           UnlockabilityStatus `json:"unlockabilityStatus"`
	SerializedAchievementTrackers json.RawMessage     `json:"serializedAchievementTrackers"`
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
	var v map[string]json.RawMessage
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
		var raw json.RawMessage
		err = json.Unmarshal(v["productionLogic"], &raw)
		if err != nil {
			return err
		}
		c.ProductionLogic = raw
	}

	return nil
}
