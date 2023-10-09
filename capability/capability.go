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

	//Always executable
	ListActions
	ListCharacterStats
	TileInfo

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
	case "listactions":
		return ListActions
	case "listcharstats":
		return ListCharacterStats
	case "tileinfo":
		return TileInfo
	default:
		return None
	}
}

func CapabilityToString(cap Capability) string {
	switch cap {
	case Forage:
		return "Forage"
	case WoodCut:
		return "WoodCut"
	case Mine:
		return "Mine"
	case Fight:
		return "Fight"
	case Move:
		return "Move"
	default:
		return "None"
	}
}
