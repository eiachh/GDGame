package capability

import "strings"

//go:generate stringer -type=Capability
// Capability represents a tool capability.
type Capability int

const (
	Forage  Capability = iota
	WoodCut            //1
	Mine               //2
	Fight

	Move

	None
	// Add more capabilities as needed
)

type MineTarget int

const (
	Copper Capability = iota
	Iron              //1
	Gold              //2
	Titan
)

func StringToCapability(cap string) Capability {
	switch strings.ToLower(cap) {
	case "forage":
		return Forage
	case "woodcut":
		return WoodCut
	case "mine":
		return Mine
	case "fight":
		return Fight
	case "move":
		return Move
	default:
		return None
	}
}
