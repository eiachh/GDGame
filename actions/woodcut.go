package action

import (
	capability "GDGame/commons/capability"
	item "GDGame/commons/item"
	tile "GDGame/map"
	"GDGame/player"
	"math"
	"strconv"
	"strings"
)

var (
	OakMinLevel           int
	OakDensityBase        int
	OakDefaultYield       float64
	OakMaxDensityModifier float64
	OakToolModifier       float64
)

func init() {

	OakMinLevel = 1
	OakDensityBase = 40
	OakDefaultYield = 1.5
	OakMaxDensityModifier = 2.0
	OakToolModifier = 0.7
	//OakToolModifBase = 0.5
}

func (acc *ActionController) executeWoodCut(player *player.Player, onTile *tile.Tile, target string) (bool, string) {
	bestTool := acc.getBestUsableTool(player.Inventory, capability.WoodCut)
	switch strings.ToLower(target) {
	case "oak":
		yield := mineOak(player, onTile, bestTool)
		return true, ("Successfully cut oak tree " + strconv.Itoa(yield) + " Oak")
	default:
		return false, "Failed to cut tree: " + target
	}
}

func mineOak(player *player.Player, onTile *tile.Tile, tool *item.Item) int {
	yield := CalculateOakYield(tool.ToolLevel, onTile.Oak)

	for i := 0; i < yield; i++ {
		Oak := item.Get(item.Oak)
		player.Inventory.AddItem(Oak)
	}

	onTile.Oak = onTile.Oak - yield

	return yield
}

func CalculateOakYield(toolLevel int, OakDensity int) int {
	OakToolModifStep := 1.0 - OakToolModifier
	FinalToolModif := 1.0

	densityModifier := float64(OakDensity) / float64(OakDensityBase)
	densityModifier = math.Min(densityModifier, OakMaxDensityModifier)

	if toolLevel < OakMinLevel {
		return 0
	}

	if toolLevel > OakMinLevel {
		FinalToolModif = OakToolModifier + (OakToolModifStep * float64(toolLevel-OakMinLevel))
	}

	return int(math.Round(float64(OakDefaultYield) * FinalToolModif * densityModifier))
}
