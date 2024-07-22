package mines

type Board struct {
	cells map[Location]*Cell
}

func NewBoard(size Size) *Board {
	return &Board{
		cells: make(map[Location]*Cell, size.Width*size.Height),
	}
}

func (board *Board) PlaceCell(location Location) {
	board.cells[location] = NewEmptyCell(location)
}

func (board *Board) PlaceMine(location Location) {
	board.cells[location] = NewMinedCell(location)
}

func (board *Board) PlaceVoid(location Location) {
	delete(board.cells, location)
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
		cell := board.Retrieve(adjacentLocation)

		if cell.Mined {
			adjacentMines++
		}
	}

	return adjacentMines
}

func (board *Board) CountUnopenSafeCells() int {
	unopenedCellsCount := 0

	for _, cell := range board.cells {
		if !cell.Mined {
			if cell.Status != Opened {
				unopenedCellsCount++
			}
		}
	}

	return unopenedCellsCount
}

func (board *Board) CountUnflaggedMines() int {
	unflaggedMinesCount := 0

	for _, cell := range board.cells {
		if cell.Mined {
			if cell.Status != Flagged {
				unflaggedMinesCount++
			}
		}
	}

	return unflaggedMinesCount
}
