package mines

type Cellmap = map[Location]*Cell

type Board struct {
	cells Cellmap
}

func newBoard(cells Cellmap) *Board {
	return &Board{
		cells,
	}
}

func (board *Board) Retrieve(location Location) *Cell {
	maybeCell, ok := board.cells[location]

	if ok {
		return maybeCell
	}

	return Void
}

func (board *Board) AdjacentMines(location Location) int {
	adjacentLocations := location.AdjacentLocations()
	adjacentMines := 0

	for _, adjacentLocation := range adjacentLocations {
		adjacentCell := board.Retrieve(adjacentLocation)

		if adjacentCell.Mined() {
			adjacentMines++
		}
	}

	return adjacentMines
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
