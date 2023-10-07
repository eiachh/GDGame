package tile

import (
	"testing"
)

func TestMapGenerate(t *testing.T) {
	myMap := GenerateMap()

	if len(myMap.Tiles) != myMap.Height || len(myMap.Tiles[0]) != myMap.Width {
		t.Errorf("Expected height: '%d', but got: '%d'", myMap.Height, len(myMap.Tiles))
	}

	if len(myMap.Tiles[0]) != myMap.Width {
		t.Errorf("Expected width: '%d', but got: '%d'", myMap.Width, len(myMap.Tiles[0]))
	}

	if myMap.Tiles[0][0] == nil || myMap.Tiles[myMap.Width-1][myMap.Height-1] == nil {
		t.Errorf("Expected map edges are nil")
	}
}
