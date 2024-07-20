package mines

type Cellstate int

const (
	Opened Cellstate = iota
	Closed
	Flagged
	Unfathomable = 0xFFFF
)

type Cell struct {
	Position Location
	Status   Cellstate
	Mined    bool
}

var Void = &Cell{
	Position: Location{-1, -1},
	Status:   Unfathomable,
	Mined:    false,
}

func NewEmptyCell(location Location) *Cell {
	return &Cell{
		Position: location,
		Status:   Closed,
		Mined:    false,
	}
}

func NewMinedCell(location Location) *Cell {
	return &Cell{
		Position: location,
		Status:   Closed,
		Mined:    true,
	}
}
