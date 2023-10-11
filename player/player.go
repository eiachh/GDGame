// player/player.go
package player

import (
	capability "GDGame/commons/capability"
	item "GDGame/commons/item"
)

var DefaultInventorySize = 20

// Player represents a player with an inventory and a name.
type Player struct {
	PlayerName string
	Inventory  *Inventory
	X          int
	Y          int
}

// NewPlayer creates a new player with the given name and maximum inventory size.
func NewPlayer(playerName string, invMax, x, y int) *Player {
	return &Player{
		PlayerName: playerName,
		Inventory:  NewInventory(invMax),
		X:          x,
		Y:          y,
	}
}

// Inventory represents a player's inventory.
type Inventory struct {
	Items  []*item.Item // Import the Item type from the "item" package
	InvMax int
}

// NewInventory creates a new inventory with the given maximum size.
func NewInventory(invMax int) *Inventory {
	return &Inventory{
		Items:  make([]*item.Item, 0), // Initialize an empty slice of items
		InvMax: invMax,
	}
}

// InitializeInventory initializes an inventory with items.
func InitializeInventory(invMax int, items []*item.Item) *Inventory {
	inv := NewInventory(invMax)
	for _, item := range items {
		inv.AddItem(item)
	}
	return inv
}

// AddItem adds an item to the inventory if there is space available.
func (inv *Inventory) AddItem(item *item.Item) bool {
	if len(inv.Items) < inv.InvMax {
		inv.Items = append(inv.Items, item)
		return true
	}
	return false // Inventory is full, item not added
}

func (inv *Inventory) AddMultipleOf(itemId item.ItemId, amount int) {
	for i := 0; i < amount; i++ {
		itemToAdd := item.Get(itemId)
		inv.AddItem(itemToAdd)
	}
}

// Returns the item from the inventory, nil if not found
func (inv *Inventory) GetItems(id int) []*item.Item {
	retItems := make([]*item.Item, 0)
	for _, item := range inv.Items {
		if item.ID == id {
			retItems = append(retItems, item)
		}
	}

	return retItems
}

func (player *Player) PlayerCapabilities() []capability.Capability {
	capabilities := make([]capability.Capability, 0, len(player.Inventory.Items))

	if player.CanPlayerMove() {
		capabilities = append(capabilities, capability.Move)
	}

	//A character can always forage if tile allows.
	capabilities = append(capabilities, capability.Forage)

	for _, item := range player.Inventory.Items {
		if item != nil && item.Capability != capability.None {
			capabilities = append(capabilities, item.Capability)
		}
	}
	return capabilities
}

// GetInventorySize returns the current size of the inventory.
func (inv *Inventory) GetInventorySize() int {
	return len(inv.Items)
}

func (player *Player) CanPlayerMove() bool {
	return true
}
