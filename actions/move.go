package action

import (
	"GDGame/player"
	"strings"
)

const (
	North int = iota
	East
	South
	West
)

func (acc *ActionController) executeMove(player *player.Player, direction int) bool {
	onTile := acc.GameMap.Tiles[player.X][player.Y]
	switch direction {
	case North:
		if onTile.Y+1 >= acc.GameMap.Height {
			return false
		}
		onTile.OnTileLeave(player)
		player.Y = player.Y + 1
		acc.GameMap.Tiles[player.X][player.Y].OnTileEnter(player)
	case East:
		if onTile.X+1 >= acc.GameMap.Width {
			return false
		}
		onTile.OnTileLeave(player)
		player.X = player.X + 1
		acc.GameMap.Tiles[player.X][player.Y].OnTileEnter(player)
	case South:
		if onTile.Y-1 < 0 {
			return false
		}
		onTile.OnTileLeave(player)
		player.Y = player.Y - 1
		acc.GameMap.Tiles[player.X][player.Y].OnTileEnter(player)
	case West:
		if onTile.X-1 < 0 {
			return false
		}
		onTile.OnTileLeave(player)
		player.X = player.X - 1
		acc.GameMap.Tiles[player.X][player.Y].OnTileEnter(player)
	}

	return true
}

func DirectionToInt(dir string) int {
	switch strings.ToLower(dir) {
	case "north":
		return North
	case "east":
		return East
	case "south":
		return South
	case "west":
		return West
	default:
		return North
	}
}
