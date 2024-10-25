package board

import (
	"math"

	"github.com/federico-paolillo/mines/pkg/dimensions"
)

type Cellstate string

const (
	Opened       Cellstate = "open"
	Closed                 = "closed"
	Flagged                = "flagged"
	Unfathomable           = "unfathomable"
)

type Cell struct {
	position      dimensions.Location
	status        Cellstate
	mined         bool
	adjacentMines int
}

var Void = &Cell{
	position: dimensions.Location{X: math.MinInt32, Y: math.MinInt32},
	status:   Unfathomable,
	mined:    false,
}

func NewSafeCell(
	location dimensions.Location,
	adjacentMines int,
) *Cell {
	return &Cell{
		position:      location,
		status:        Closed,
		mined:         false,
		adjacentMines: adjacentMines,
	}
}

func NewMineCell(
	location dimensions.Location,
	adjacentMines int,
) *Cell {
	return &Cell{
		position:      location,
		status:        Closed,
		mined:         true,
		adjacentMines: adjacentMines,
	}
}

func (cell *Cell) Open() {
	if cell.HasStatus(Closed, Flagged) {
		cell.status = Opened
	}
}

func (cell *Cell) Flag() {
	if cell.HasStatus(Closed) {
		cell.status = Flagged
	}
}

func (cell *Cell) HasStatus(statuses ...Cellstate) bool {
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

func (cell *Cell) Position() dimensions.Location {
	return cell.position
}

func (cell *Cell) AdjacentMines() int {
	return cell.adjacentMines
}

func (cell *Cell) Status() Cellstate {
	return cell.status
}
