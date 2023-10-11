package action

import (
	item "GDGame/commons/item"
	tile "GDGame/map"
	"GDGame/player"
	"testing"
)

func TestOakYieldBase(t *testing.T) {
	toolLevel := 2
	oakDensity := 40
	yield := CalculateOakYield(toolLevel, oakDensity)
	if yield != 2 {
		t.Errorf("Expected copper yield 2 but got '%d'", yield)
	}
}

func TestOakYieldMinTool(t *testing.T) {
	toolLevel := 1
	oakDensity := 40
	yield := CalculateOakYield(toolLevel, oakDensity)
	if yield != 2 {
		t.Errorf("Expected copper yield 2 but got '%d'", yield)
	}
}

func TestOakYieldGoodTool(t *testing.T) {
	toolLevel := 3
	oakDensity := 10
	yield := CalculateCopperYield(toolLevel, oakDensity)
	if yield != 3 {
		t.Errorf("Expected copper yield 3 but got '%d'", yield)
	}
}

func TestOakYielDoubleDensity(t *testing.T) {
	toolLevel := 2
	oakDensity := 80
	yield := CalculateCopperYield(toolLevel, oakDensity)
	if yield != 4 {
		t.Errorf("Expected copper yield 2 but got '%d'", yield)
	}
}

// With toolLvl 2 ore density 10 expected result 2 copper in inventory and -2 to the density
func TestPlayerCutOak(t *testing.T) {
	oakDensity := 40
	cuttingTile := tile.NewTile(0, 0, oakDensity, 0, 0, 0, 0)

	axe2 := item.Get(item.Axe2)
	player1 := player.NewPlayer("testplayer", 10, 0, 0)
	player1.Inventory.AddItem(axe2)

	ac := NewActionController(tile.NewMap())

	ac.executeWoodCut(player1, cuttingTile, "oak")

	oakCountInInventory := len(player1.Inventory.GetItems(int(item.Oak)))
	if oakCountInInventory != 2 {
		t.Errorf("Expected copper count in inventory 2 but got: '%d'", oakCountInInventory)
	}

	if cuttingTile.Oak != oakDensity-oakCountInInventory {
		t.Errorf("Expected leftover copper on tile 8 but has '%d'", cuttingTile.Copper)
	}
}
