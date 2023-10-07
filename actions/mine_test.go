package action

import (
	"GDGame/item"
	tile "GDGame/map"
	"GDGame/player"
	"testing"
)

func TestCopperYieldBase(t *testing.T) {
	CopperMinLevel = 1
	CopperDensityBase = 10
	CopperDefaultYield = 2

	toolLevel := 2
	copperDensity := 10
	yield := CalculateCopperYield(toolLevel, copperDensity)
	if yield != 2 {
		t.Errorf("Expected copper yield 2 but got '%d'", yield)
	}
}

func TestCopperYieldMinTool(t *testing.T) {
	CopperMinLevel = 1
	CopperDensityBase = 10
	CopperDefaultYield = 2

	toolLevel := 1
	copperDensity := 10
	yield := CalculateCopperYield(toolLevel, copperDensity)
	if yield != 2 {
		t.Errorf("Expected copper yield 2 but got '%d'", yield)
	}
}

func TestCopperYieldGoodTool(t *testing.T) {
	CopperMinLevel = 1
	CopperDensityBase = 10
	CopperDefaultYield = 2

	toolLevel := 3
	copperDensity := 10
	yield := CalculateCopperYield(toolLevel, copperDensity)
	if yield != 3 {
		t.Errorf("Expected copper yield 3 but got '%d'", yield)
	}
}

func TestCopperYielDoubleDensity(t *testing.T) {
	CopperMinLevel = 1
	CopperDensityBase = 10
	CopperDefaultYield = 2

	toolLevel := 2
	copperDensity := 20
	yield := CalculateCopperYield(toolLevel, copperDensity)
	if yield != 4 {
		t.Errorf("Expected copper yield 2 but got '%d'", yield)
	}
}

// With toolLvl 2 ore density 10 expected result 2 copper in inventory and -2 to the density
func TestPlayerMineCopper(t *testing.T) {
	copperDensity := 10
	miningTile := tile.NewTile(0, 0, 0, copperDensity, 0, 0, 0)

	pickaxe2 := item.Get(item.Pickaxe2)
	player1 := player.NewPlayer("testplayer", 10, 0, 0)
	player1.Inventory.AddItem(pickaxe2)

	ac := NewActionController(tile.NewMap())

	ac.executeMine(player1, miningTile, "copper")

	copperCountInInventory := len(player1.Inventory.GetItems(int(item.Copper)))
	if copperCountInInventory != 2 {
		t.Errorf("Expected copper count in inventory 2 but got: '%d'", copperCountInInventory)
	}

	if miningTile.Copper != 8 {
		t.Errorf("Expected leftover copper on tile 8 but has '%d'", miningTile.Copper)
	}
}
