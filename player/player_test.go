package player

import (
	item "GDGame/commons/item"
	"testing"
)

func TestNewPlayer(t *testing.T) {

	player1 := NewPlayer("Bob", 64, 0, 0)

	if player1.PlayerName != "Bob" {
		t.Errorf("Expected player name Bob, but got '%s'", player1.PlayerName)
	}

	if player1.Inventory.InvMax != 64 {
		t.Errorf("Expected inventory size 64, but got '%d'", player1.Inventory.InvMax)
	}
}

func TestSingleItemAddAndGet(t *testing.T) {

	player1 := NewPlayer("Bob", 64, 0, 0)

	adminpickaxe := item.Get(item.ADMINpickaxe)
	player1.Inventory.AddItem(adminpickaxe)
	if len(player1.Inventory.GetItems(adminpickaxe.ID)) == 0 {
		t.Errorf("Expected to find item with ID 4 in inventory but got empty list")
	}

	if len(player1.Inventory.GetItems(69)) != 0 {
		t.Errorf("Expected to get empty list but len: '%d'", len(player1.Inventory.GetItems(69)))
	}
}

func TestFullInventory(t *testing.T) {
	player1 := NewPlayer("Bob", 2, 0, 0)

	adminpickaxe := item.Get(item.ADMINpickaxe)
	player1.Inventory.AddItem(adminpickaxe)

	if len(player1.Inventory.GetItems(adminpickaxe.ID)) != 1 {
		t.Errorf("Expected to find item with ID 4 in inventory but got empty list")
	}

	adminpickaxe2 := item.Get(item.ADMINpickaxe)
	player1.Inventory.AddItem(adminpickaxe2)

	if len(player1.Inventory.GetItems(adminpickaxe2.ID)) != 2 {
		t.Errorf("Expected to find 2 item with ID 2 but got len: '%d'", len(player1.Inventory.GetItems(4)))
	}

	adminpickaxe3 := item.Get(item.ADMINpickaxe)
	if player1.Inventory.AddItem(adminpickaxe3) {
		t.Errorf("Expected add item to fail for full inventory but got true")
	}
}
