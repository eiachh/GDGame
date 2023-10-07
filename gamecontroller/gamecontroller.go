package gamecontroller

import (
	action "GDGame/actions"
	"GDGame/capability"
	tile "GDGame/map"
	"GDGame/player"
)

type BasicCommand struct {
	OwnerId  string `json:"OwnerId"`
	Command  string `json:"Command"`
	ExtraArg string `json:"ExtraArg"`
}

type RegisterCommand struct {
	OwnerId    string `json:"OwnerId"`
	PlayerName string `json:"PlayerName"`
}

var GameMap *tile.Map
var AcController *action.ActionController

func init() {
	GameMap = tile.GenerateMap()
	AcController = action.NewActionController(GameMap)
}

func RegisterPlayer(regCommand RegisterCommand) (bool, *player.Player) {
	regPlayer := player.NewPlayer(regCommand.PlayerName, player.DefaultInventorySize, GameMap.SpawnX, GameMap.SpawnY)
	return (player.RegisterPlayer(regCommand.OwnerId, regPlayer)), regPlayer
}

func HandleBasicCommand(bcomm BasicCommand) {
	target := player.GetPlayerByOwnerId(bcomm.OwnerId)
	AcController.ExecuteAction(
		target,
		GameMap.Tiles[target.X][target.Y],
		capability.StringToCapability(bcomm.Command),
		bcomm.ExtraArg)
}
