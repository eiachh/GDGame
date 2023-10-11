package gamecontroller

import (
	action "GDGame/actions"
	"GDGame/commons"
	capability "GDGame/commons/capability"
	tile "GDGame/map"
	"GDGame/player"
	"fmt"
	"reflect"
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
	case capability.TileInfo:
		return TileInfo(*tile)
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
	response = strings.TrimSuffix(response, " [")
	return response
}

func TileInfo(targetTile tile.Tile) string {
	var response string
	v := reflect.ValueOf(targetTile)

	// Ensure v is a struct; if not, return
	if v.Kind() != reflect.Struct {
		return "Error"
	}

	// Iterate through the struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)
		response += fmt.Sprintf("%s: %v\n", field.Name, value.Interface())
	}

	return response
}
