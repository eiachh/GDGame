package action

import (
	capability "GDGame/commons/capability"
	tile "GDGame/map"
	"GDGame/player"
	"testing"
)

func TestSimpleMove(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, gameMap.SpawnX, gameMap.SpawnY)

	ac := NewActionController(gameMap)

	success, _ := ac.executeMove(player1, North)
	if !success {
		t.Errorf("Player1 should be able to move north but was blocked")
	}
	if player1.Y != gameMap.SpawnY+1 {
		t.Errorf("Player1 is Y expected 501 but got: '%d'", player1.Y)
	}

	success, _ = ac.executeMove(player1, East)
	if !success {
		t.Errorf("Player1 should be able to move east but was blocked")
	}
	if player1.X != gameMap.SpawnX+1 {
		t.Errorf("Player1 is X expected 501 but got: '%d'", player1.X)
	}

	success, _ = ac.executeMove(player1, South)
	if !success {
		t.Errorf("Player1 should be able to move south but was blocked")
	}
	if player1.Y != gameMap.SpawnY {
		t.Errorf("Player1 is Y expected 500 but got: '%d'", player1.Y)
	}

	success, _ = ac.executeMove(player1, West)
	if !success {
		t.Errorf("Player1 should be able to move west but was blocked")
	}
	if player1.X != gameMap.SpawnX {
		t.Errorf("Player1 is X expected 500 but got: '%d'", player1.X)
	}

}

func TestOutOfBoundsNorth(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, gameMap.SpawnX, gameMap.Height-1)

	ac := NewActionController(gameMap)

	success, _ := ac.executeMove(player1, North)
	if success {
		t.Errorf("Player1 should NOT be able to move north but it could")
	}
}

func TestOutOfBoundsEast(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, gameMap.Width-1, gameMap.SpawnY)

	ac := NewActionController(gameMap)

	success, _ := ac.executeMove(player1, East)
	if success {
		t.Errorf("Player1 should NOT be able to move east but it could")
	}
}

func TestOutOfBoundsSouth(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, gameMap.SpawnX, 0)

	ac := NewActionController(gameMap)

	success, _ := ac.executeMove(player1, South)
	if success {
		t.Errorf("Player1 should NOT be able to move south but it could")
	}
}

func TestOutOfBoundsWest(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, 0, gameMap.SpawnY)

	ac := NewActionController(gameMap)

	success, _ := ac.executeMove(player1, West)
	if success {
		t.Errorf("Player1 should NOT be able to move west but it could")
	}
}

func TestExecAction(t *testing.T) {
	gameMap := tile.GenerateMap()
	player1 := player.NewPlayer("testPlayer", 10, gameMap.SpawnX, gameMap.SpawnY)
	ac := NewActionController(gameMap)

	success, _ := ac.ExecuteAction(player1, gameMap.Tiles[player1.X][player1.Y], capability.Move, "North")
	if !success {
		t.Errorf("Player1 should be able to move north but was blocked")
	}
	if player1.Y != gameMap.SpawnY+1 {
		t.Errorf("Player1 is Y expected 501 but got: '%d'", player1.Y)
	}

	success, _ = ac.ExecuteAction(player1, gameMap.Tiles[player1.X][player1.Y], capability.Move, "East")
	if !success {
		t.Errorf("Player1 should be able to move east but was blocked")
	}
	if player1.X != gameMap.SpawnX+1 {
		t.Errorf("Player1 is X expected 501 but got: '%d'", player1.X)
	}

	success, _ = ac.ExecuteAction(player1, gameMap.Tiles[player1.X][player1.Y], capability.Move, "South")
	if !success {
		t.Errorf("Player1 should be able to move south but was blocked")
	}
	if player1.Y != gameMap.SpawnY {
		t.Errorf("Player1 is Y expected 500 but got: '%d'", player1.Y)
	}

	success, _ = ac.ExecuteAction(player1, gameMap.Tiles[player1.X][player1.Y], capability.Move, "West")
	if !success {
		t.Errorf("Player1 should be able to move west but was blocked")
	}
	if player1.X != gameMap.SpawnX {
		t.Errorf("Player1 is X expected 500 but got: '%d'", player1.X)
	}
}
