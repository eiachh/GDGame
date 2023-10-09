package action

import (
	"GDGame/capability"
	"GDGame/item"
	tile "GDGame/map"
	"GDGame/player"
	"math"
	"strconv"
	"strings"
)

var (
	MiningToolIds []int

	CopperMinLevel           int
	CopperDensityBase        int
	CopperDefaultYield       int
	CopperMaxDensityModifier float64
	CopperToolModifier       float64
)

func init() {
	MiningToolIds = []int{3, 4}

	CopperMinLevel = 1
	CopperDensityBase = 10
	CopperDefaultYield = 2
	CopperMaxDensityModifier = 2.0
	CopperToolModifier = 0.7
	//CopperToolModifBase = 0.5
}

func (acc *ActionController) executeMine(player *player.Player, onTile *tile.Tile, target string) (bool, string) {
	bestTool := acc.getBestUsableTool(player.Inventory, capability.Mine)
	switch strings.ToLower(target) {
	case "copper":
		yield := mineCopper(player, onTile, bestTool)
		return true, ("Successfully mined " + strconv.Itoa(yield) + " copper")
	default:
		return false, "Failed to mine ore: " + target
	}
}

func mineCopper(player *player.Player, onTile *tile.Tile, tool *item.Item) int {
	yield := CalculateCopperYield(tool.ToolLevel, onTile.Copper)

	for i := 0; i < yield; i++ {
		copper := item.Get(item.Copper)
		player.Inventory.AddItem(copper)
	}

	onTile.Copper = onTile.Copper - yield

	return yield
}

func CalculateCopperYield(toolLevel int, copperDensity int) int {
	CopperToolModifStep := 1.0 - CopperToolModifier
	FinalToolModif := 1.0

	densityModifier := float64(copperDensity) / float64(CopperDensityBase)
	densityModifier = math.Min(densityModifier, CopperMaxDensityModifier)

	if toolLevel < CopperMinLevel {
		return 0
	}

	if toolLevel > CopperMinLevel {
		FinalToolModif = CopperToolModifier + (CopperToolModifStep * float64(toolLevel-CopperMinLevel))
	}

	return int(math.Round(float64(CopperDefaultYield) * FinalToolModif * densityModifier))
}
