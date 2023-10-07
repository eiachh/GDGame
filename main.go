package main

import (
	"GDGame/rest"
)

func main() {
	rest.StartRestApi(8080)

	/* player1 := player.NewPlayer("Bob", 64, 0, 0)
	wood := item.Get(1)
	axe := item.Get(2)
	pickaxe := item.Get(3)
	adminpickaxe := item.Get(4)
	player1.Inventory.AddItem(wood)
	player1.Inventory.AddItem(axe)
	player1.Inventory.AddItem(pickaxe)
	player1.Inventory.AddItem(adminpickaxe)
	fmt.Printf("Player name: %s\n", player1.PlayerName)
	fmt.Printf("Item 1 name: %s\n", player1.Inventory.Items[1].Name)

	//randomTile := tile.Forest

	randomTile := tile.TileType.GetBaseTile(tile.Forest)
	fmt.Printf("%d", randomTile.Wood)

	for _, i := range randomTile.TileCapabilities() {
		fmt.Printf("Capability of randomTile: %d\n", i)
	}

	fmt.Println() */
}
