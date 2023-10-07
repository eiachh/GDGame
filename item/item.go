// item/item.go
package item

import "GDGame/capability"

type ItemId int

const (
	Wood ItemId = iota
	Copper

	Axe = iota + 500
	Pickaxe
	Pickaxe2
	ADMINpickaxe

	Empty
	// Add more capabilities as needed
)

// Item represents an item with a name and quantity.
type Item struct {
	ID          int
	Name        string
	Capability  capability.Capability
	ToolLevel   int
	Description string
	Weight      int
}

func Get(id ItemId) *Item {
	name, capability, toolLevel, desc, weight := querryItemData(id)
	return &Item{
		ID:          int(id),
		Name:        name,
		Capability:  capability,
		ToolLevel:   toolLevel,
		Description: desc,
		Weight:      weight,
	}
}

func querryItemData(id ItemId) (string, capability.Capability, int, string, int) {
	//name, capability, toolLevel, desc, weight := querryItemData(id)
	switch id {
	case Wood:
		return "wood", capability.None, 0, "somedesc", 1
	case Copper:
		return "copper", capability.None, 0, "somedesc", 1
	case Axe:
		return "axe", capability.WoodCut, 1, "somedesc", 1
	case Pickaxe:
		return "pickaxe", capability.Mine, 1, "somedesc", 1
	case Pickaxe2:
		return "pickaxe2", capability.Mine, 2, "somedesc", 1
	case ADMINpickaxe:
		return "ADMINpickaxe", capability.Mine, 99, "somedesc", 1
	default:
		return "unknown", capability.None, 0, "unknown", 1
	}
}
