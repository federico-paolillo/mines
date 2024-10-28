package board

import "github.com/federico-paolillo/mines/pkg/dimensions"

type Cellmap = map[dimensions.Location]*Cell

type Board struct {
	size  dimensions.Size
	cells Cellmap
}

func newBoard(size dimensions.Size, cells Cellmap) *Board {
	return &Board{
		size,
		cells,
	}
}

func (board *Board) Retrieve(location dimensions.Location) *Cell {
	maybeCell, ok := board.cells[location]

	if ok {
		return maybeCell
	}

	return Void
}

func (board *Board) CountUnopenSafeCells() int {
	unopenedCellsCount := 0

	for _, cell := range board.cells {
		if cell.Safe() {
			if cell.HasStatus(ClosedCell, FlaggedCell) {
				unopenedCellsCount++
			}
		}
	}

	return unopenedCellsCount
}

func (board *Board) Size() dimensions.Size {
	return board.size
}

func (board *Board) CountAdjacentCellsOfStatus(status Cellstate, location dimensions.Location) int {
	cells := board.RetrieveAdjacentCellsOfStatus(status, location)

	count := len(cells)

	return count
}

func (board *Board) RetrieveAdjacentCellsOfStatus(status Cellstate, location dimensions.Location) []*Cell {
	cells := make([]*Cell, 0, 8)

	for _, location := range location.AdjacentLocations() {
		cell := board.Retrieve(location)

		if cell.HasStatus(status) {
			cells = append(cells, cell)
		}
	}

	return cells
}
