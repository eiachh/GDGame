package gamecontroller

import (
	action "GDGame/actions"
	"GDGame/capability"
	"GDGame/commons"
	tile "GDGame/map"
	"GDGame/player"
	"strconv"
	"strings"
)

var GameMap *tile.Map
var AcController *action.ActionController

func init() {
	GameMap = tile.GenerateMap()
	AcController = action.NewActionController(GameMap)
}

func RegisterPlayer(regCommand commons.RegisterCommand) (bool, string, *player.Player) {
	regPlayer := player.NewPlayer(regCommand.PlayerName, player.DefaultInventorySize, GameMap.SpawnX, GameMap.SpawnY)
	succ, response := player.RegisterPlayer(regCommand.OwnerId, regPlayer)
	return succ, response, regPlayer
}

func HandleBasicCommand(bcomm commons.BasicCommand) string {
	target := player.GetPlayerByOwnerId(bcomm.OwnerId)

	if target == nil {
		return "Player not registered/found!"
	}

	_, response := AcController.ExecuteAction(
		target,
		GameMap.Tiles[target.X][target.Y],
		capability.StringToCapability(bcomm.Command),
		bcomm.ExtraArg)
	return response
}

func HandleInfoCommand(bcomm commons.BasicCommand) string {
	targetPlayer := player.GetPlayerByOwnerId(bcomm.OwnerId)

	if targetPlayer == nil {
		return "Player not registered/found!"
	}

	tile := GameMap.Tiles[targetPlayer.X][targetPlayer.Y]

	switch capability.StringToCapability(bcomm.Command) {
	case capability.ListActions:
		return ListActions(targetPlayer, tile)
	case capability.ListCharacterStats:
		return ListCharacterStats(targetPlayer)
	}
	return "Action not found (gc H.I.C.)!"
}

func ListActions(targetPlayer *player.Player, tile *tile.Tile) string {
	capabilities := AcController.GetPossibleActions(targetPlayer, tile)
	response := "Your player can: "

	for _, capAsInt := range capabilities {
		response += capability.CapabilityToString(capability.Capability(capAsInt)) + ", "
	}
	response = strings.TrimSuffix(response, ", ")
	return response
}

func ListCharacterStats(targetPlayer *player.Player) string {
	response := "Char: " + targetPlayer.PlayerName +
		"\n Location[X:" + strconv.Itoa(targetPlayer.X) +
		" ][Y:" + strconv.Itoa(targetPlayer.Y) + " ]"

	if len(targetPlayer.Inventory.Items) == 0 {
		response += "\n Items: [NONE YET]"
	} else {
		response += "\n Items: ["
	}

	for _, item := range targetPlayer.Inventory.Items {
		response += item.Name + "] ["
	}
	response = strings.TrimSuffix(response, "] [")
	return response
}
