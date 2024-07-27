package board

import "github.com/federico-paolillo/mines/pkg/dimensions"

type Cellmap = map[dimensions.Location]*Cell

type Board struct {
	cells Cellmap
}

func newBoard(cells Cellmap) *Board {
	return &Board{
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
			if cell.Status(Closed, Flagged) {
				unopenedCellsCount++
			}
		}
	}

	return unopenedCellsCount
}

func (board *Board) CountUnflaggedMines() int {
	unflaggedMinesCount := 0

	for _, cell := range board.cells {
		if cell.Mined() {
			if cell.Status(Closed) {
				unflaggedMinesCount++
			}
		}
	}

	return unflaggedMinesCount
}
