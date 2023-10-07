package action

import (
	"GDGame/capability"
	tile "GDGame/map"
	"GDGame/player"
	"testing"
)

func TestSimpleMove(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, 500, 500)

	ac := NewActionController(gameMap)

	if !ac.executeMove(player1, North) {
		t.Errorf("Player1 should be able to move north but was blocked")
	}
	if player1.Y != 501 {
		t.Errorf("Player1 is Y expected 501 but got: '%d'", player1.Y)
	}

	if !ac.executeMove(player1, East) {
		t.Errorf("Player1 should be able to move east but was blocked")
	}
	if player1.X != 501 {
		t.Errorf("Player1 is X expected 501 but got: '%d'", player1.X)
	}

	if !ac.executeMove(player1, South) {
		t.Errorf("Player1 should be able to move south but was blocked")
	}
	if player1.Y != 500 {
		t.Errorf("Player1 is Y expected 500 but got: '%d'", player1.Y)
	}

	if !ac.executeMove(player1, West) {
		t.Errorf("Player1 should be able to move west but was blocked")
	}
	if player1.X != 500 {
		t.Errorf("Player1 is X expected 500 but got: '%d'", player1.X)
	}

}

func TestOutOfBoundsNorth(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, 500, gameMap.Height-1)

	ac := NewActionController(gameMap)

	if ac.executeMove(player1, North) {
		t.Errorf("Player1 should NOT be able to move north but it could")
	}
}

func TestOutOfBoundsEast(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, gameMap.Width-1, 500)

	ac := NewActionController(gameMap)

	if ac.executeMove(player1, East) {
		t.Errorf("Player1 should NOT be able to move east but it could")
	}
}

func TestOutOfBoundsSouth(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, 500, 0)

	ac := NewActionController(gameMap)

	if ac.executeMove(player1, South) {
		t.Errorf("Player1 should NOT be able to move south but it could")
	}
}

func TestOutOfBoundsWest(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, 0, 500)

	ac := NewActionController(gameMap)

	if ac.executeMove(player1, West) {
		t.Errorf("Player1 should NOT be able to move west but it could")
	}
}

func TestExecAction(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, 500, 500)
	ac := NewActionController(gameMap)

	if !ac.ExecuteAction(player1, gameMap.Tiles[player1.X][player1.Y], capability.Move, "North") {
		t.Errorf("Player1 should be able to move north but was blocked")
	}
	if player1.Y != 501 {
		t.Errorf("Player1 is Y expected 501 but got: '%d'", player1.Y)
	}

	if !ac.ExecuteAction(player1, gameMap.Tiles[player1.X][player1.Y], capability.Move, "East") {
		t.Errorf("Player1 should be able to move east but was blocked")
	}
	if player1.X != 501 {
		t.Errorf("Player1 is X expected 501 but got: '%d'", player1.X)
	}

	if !ac.ExecuteAction(player1, gameMap.Tiles[player1.X][player1.Y], capability.Move, "South") {
		t.Errorf("Player1 should be able to move south but was blocked")
	}
	if player1.Y != 500 {
		t.Errorf("Player1 is Y expected 500 but got: '%d'", player1.Y)
	}

	if !ac.ExecuteAction(player1, gameMap.Tiles[player1.X][player1.Y], capability.Move, "West") {
		t.Errorf("Player1 should be able to move west but was blocked")
	}
	if player1.X != 500 {
		t.Errorf("Player1 is X expected 500 but got: '%d'", player1.X)
	}
}
