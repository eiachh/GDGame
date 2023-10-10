package action

import (
	"GDGame/item"
	tile "GDGame/map"
	"GDGame/player"
	"math"
	"strconv"
)

var (
	ForageDensityBase        = 15
	ForageDefaultYield       = 3
	ForageMaxDensityModifier = 5.0
	ForageToolModifier       = 0.7
)

func (acc *ActionController) executeForage(player *player.Player, onTile *tile.Tile, target string) (bool, string) {
	toolLevel := 1
	ForageToolModifStep := 1.0 - ForageToolModifier
	FinalToolModif := 1.0

	densityModifier := float64(onTile.Forage) / float64(ForageDensityBase)
	densityModifier = math.Min(densityModifier, ForageMaxDensityModifier)

	FinalToolModif = ForageToolModifier + (ForageToolModifStep * float64(toolLevel))

	yield := int(math.Round(float64(ForageDefaultYield) * FinalToolModif * densityModifier))

	onTile.Forage -= yield

	player.Inventory.AddMultipleOf(item.Berry, yield)
	return true, "Foraged " + strconv.Itoa(yield) + " berries"
}
