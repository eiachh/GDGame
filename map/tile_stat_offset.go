package tile

func (tileType TileType) GetBaseTile() *Tile {

	switch tileType {
	case Plains:
		return NewTile(0, 0, 20, 5, 25, 0, 0)
	case Forest:
		return NewTile(0, 0, 60, 10, 15, 0, 0)
	default:
		return NewTile(0, 0, 0, 0, 0, 0, 0)
	}
}
