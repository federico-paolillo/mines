package mines

import "math"

type Cellstate int

const (
	Opened Cellstate = iota
	Closed
	Flagged
	Unfathomable = 0xFFFF
)

type Cell struct {
	position Location
	status   Cellstate
	mined    bool
}

var Void = &Cell{
	position: Location{math.MinInt32, math.MinInt32},
	status:   Unfathomable,
	mined:    false,
}

func NewEmptyCell(location Location) *Cell {
	return &Cell{
		position: location,
		status:   Closed,
		mined:    false,
	}
}

func NewMinedCell(location Location) *Cell {
	return &Cell{
		position: location,
		status:   Closed,
		mined:    true,
	}
}

func (cell *Cell) Open() {
	if cell.Status(Closed, Flagged) {
		cell.status = Opened
	}
}

func (cell *Cell) Flag() {
	if cell.Status(Closed) {
		cell.status = Flagged
	}
}

func (cell *Cell) Status(statuses ...Cellstate) bool {
	for _, state := range statuses {
		if cell.status == state {
			return true
		}
	}

	return false
}

func (cell *Cell) Mined() bool {
	return cell.mined == true
}

func (cell *Cell) Safe() bool {
	return cell.mined == false
}

func (cell *Cell) Position() Location {
	return cell.position
}
