package tile

import (
	capability "GDGame/commons/capability"
	"GDGame/player"
	"fmt"
)

type TileType int

const (
	Plains TileType = iota
	Desert
	Forest
	Mountain

	City
	// Add more capabilities as needed
)

type Tile struct {
	X int
	Y int

	Oak    int
	Copper int
	Forage int

	HostileLvl  int
	WildlifeLvl int
	//Hostile
	//Wildlife

	Unique string
}

func NewTile(x, y, oak, copper, forage, hostilelvl, wildlifelvl int) *Tile {
	return &Tile{
		X:           x,
		Y:           y,
		Oak:         oak,
		Copper:      copper,
		Forage:      forage,
		HostileLvl:  hostilelvl,
		WildlifeLvl: wildlifelvl,
	}
}

func (tile *Tile) SetTileAsUnique(unique string) {
	tile.Unique = unique
}

func (tile *Tile) TileCapabilities() []capability.Capability {
	capabilities := make([]capability.Capability, 0, capability.None)

	if tile.IsPlayerMoveAllowed() {
		capabilities = append(capabilities, capability.Move)
	}

	if tile.Oak > 0 {
		capabilities = append(capabilities, capability.WoodCut)
	}
	if tile.Copper > 0 {
		capabilities = append(capabilities, capability.Mine)
	}
	if tile.Forage > 0 {
		capabilities = append(capabilities, capability.Forage)
	}

	if len(capabilities) == 0 {
		capabilities = append(capabilities, capability.None)
	}
	return capabilities
}

func (tile *Tile) OnTileLeave(player *player.Player) {
	fmt.Printf("Player '%s' left tile:['%d']['%d'] ", player.PlayerName, tile.X, tile.Y)
}

func (tile *Tile) OnTileEnter(player *player.Player) {
	fmt.Printf("Player '%s' left tile:['%d']['%d'] ", player.PlayerName, tile.X, tile.Y)
}

func (tile *Tile) IsPlayerMoveAllowed() bool {
	return true
}
