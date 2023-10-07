package action

import (
	"GDGame/capability"
	"GDGame/item"
	tile "GDGame/map"
	"GDGame/player"
)

type ActionController struct {
	GameMap *tile.Map
}

func NewActionController(gameMap *tile.Map) *ActionController {
	return &ActionController{
		GameMap: gameMap,
	}
}

func (acc *ActionController) GetPossibleActions(player *player.Player, onTile *tile.Tile) []capability.Capability {

	playerCap := player.PlayerCapabilities()
	tileCap := onTile.TileCapabilities()

	possCap := make([]capability.Capability, 0, capability.None)
	for _, pCap := range playerCap {
		for _, tCap := range tileCap {
			if pCap == tCap {
				possCap = append(possCap, pCap)
			}
		}
	}

	return acc.FilterOutActions(player, onTile, possCap)
}

// Remove actions that are restricted by special factors.
func (acc *ActionController) FilterOutActions(player *player.Player, onTile *tile.Tile, capabilities []capability.Capability) []capability.Capability {
	return capabilities
}

func (acc *ActionController) CanExecuteAction(player *player.Player, onTile *tile.Tile, action capability.Capability) bool {
	playerCap := player.PlayerCapabilities()
	tileCap := onTile.TileCapabilities()

	for _, pCap := range playerCap {
		for _, tCap := range tileCap {
			if (pCap == tCap) && pCap == action {
				return true
			}
		}
	}
	return false
}

func (acc *ActionController) ExecuteAction(player *player.Player, onTile *tile.Tile, action capability.Capability, capExtraArg string) bool {
	if !acc.CanExecuteAction(player, onTile, action) {
		return false
	}

	switch action {
	case capability.Mine:
		return acc.executeMine(player, onTile, capExtraArg)
	case capability.Move:
		return acc.executeMove(player, DirectionToInt(capExtraArg))
	default:
		return false
	}
}

func (acc *ActionController) getBestUsableTool(inv *player.Inventory, cap capability.Capability) *item.Item {
	maxToolLvl := 0
	retItem := item.Get(-1)
	for _, item := range inv.Items {
		if item.Capability == cap && item.ToolLevel > maxToolLvl {
			maxToolLvl = item.ToolLevel
			retItem = item
		}
	}
	return retItem
}
