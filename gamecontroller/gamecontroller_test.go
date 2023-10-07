package gamecontroller

import (
	"GDGame/player"
	"testing"
)

func TestPlayerRegistration(t *testing.T) {
	ownerId := "ownerdid"
	regCommand := &RegisterCommand{
		OwnerId:    ownerId,
		PlayerName: "testplayer",
	}
	success, _ := RegisterPlayer(*regCommand)
	if !success {
		t.Errorf("Register player failed!")
	}

	playerByoId := player.GetPlayerByOwnerId(ownerId)
	if playerByoId == nil {
		t.Errorf("Player was not found by oId")
	}
}

func TestBasicCommand(t *testing.T) {
	ownerId := "ownerdid"

	regCommand := &RegisterCommand{
		OwnerId:    ownerId,
		PlayerName: "testplayer",
	}
	_, regPlayer := RegisterPlayer(*regCommand)
	playerYBeforeMove := regPlayer.Y
	bcommand := &BasicCommand{
		OwnerId:  ownerId,
		Command:  "Move",
		ExtraArg: "North",
	}
	HandleBasicCommand(*bcommand)

	if playerYBeforeMove >= regPlayer.Y {
		t.Errorf("Player new Y location is not higher. Old loc:'%d', Current loc: '%d'", playerYBeforeMove, regPlayer.Y)
	}
}
