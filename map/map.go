package tile

type Map struct {
	Tiles  map[int]map[int]*Tile
	Width  int
	Height int
	SpawnX int
	SpawnY int
}

func NewMap() *Map {
	width := 500
	height := 500
	return &Map{
		Tiles:  make(map[int]map[int]*Tile),
		Width:  width,
		Height: height,
		SpawnX: (width / 2),
		SpawnY: (height / 2),
	}
}

func GenerateMap() *Map {
	gameMap := NewMap()
	for row := 0; row < gameMap.Height; row++ {
		gameMap.Tiles[row] = make(map[int]*Tile)
		for col := 0; col < gameMap.Width; col++ {
			currTile := TileType.GetBaseTile(Forest)
			currTile.X = row
			currTile.Y = col
			gameMap.Tiles[row][col] = currTile
		}
	}

	gameMap.Tiles[gameMap.Width/2][gameMap.Height/2] =
		&Tile{
			X:           (gameMap.Width / 2),
			Y:           (gameMap.Height / 2),
			Oak:         0,
			Copper:      0,
			Forage:      0,
			HostileLvl:  0,
			WildlifeLvl: 0,
		}
	gameMap.Tiles[gameMap.Width/2][gameMap.Height/2].SetTileAsUnique("City")

	return gameMap
}

func (m *Map) SetTileAt(x, y int, target *Tile) {
	if m.Tiles == nil {
		NewMap()
	}
	m.Tiles[x][y] = target
}
