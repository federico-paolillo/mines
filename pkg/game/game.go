package game

import (
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

type Status = string

const (
	Playing Status = "playing"
	Lost           = "lost"
	Won            = "won"
)

type Game struct {
	board *board.Board
	lives int
}

func NewGame(lives int, board *board.Board) *Game {
	return &Game{
		board: board,
		lives: lives,
	}
}

func (game *Game) Flag(x, y int) {
	if game.ended() {
		return
	}

	cell := game.board.Retrieve(dimensions.Location{X: x, Y: y})

	if cell == board.Void {
		return
	}

	cell.Flag()
}

func (game *Game) Open(x, y int) {
	if game.ended() {
		return
	}

	cell := game.board.Retrieve(dimensions.Location{X: x, Y: y})

	if cell == board.Void {
		return
	}

	cell.Open()

	if cell.Mined() {
		game.lives--
		return
	}

	if cell.AdjacentMines() == 0 {
		game.tryCascade(cell.Position())
	}
}

// Attempt chording at specified location. To chord, the selected cell must already be open and
// must have at least one adjacent mine.
//
// When the number of adjacent mines is equal to the number of adjacent flagged cells,
// all adjacent non-flagged unopened cells will be opened.
func (game *Game) Chord(x, y int) {
	if game.ended() {
		return
	}

	location := dimensions.Location{X: x, Y: y}

	originCell := game.board.Retrieve(location)

	if originCell == board.Void {
		return
	}

	if !originCell.Status(board.Opened) {
		return
	}

	adjacentFlaggedCells := game.board.CountAdjacentCellsOfStatus(board.Flagged, location)

	if adjacentFlaggedCells != originCell.AdjacentMines() {
		return
	}

	adjacentClosedCells := game.board.RetrieveAdjacentCellsOfStatus(board.Closed, location)

	for _, closedCell := range adjacentClosedCells {
		cellLocation := closedCell.Position()
		game.Open(cellLocation.X, cellLocation.Y)
	}
}

func (game *Game) Status() Status {
	if game.lives < 0 {
		return Lost
	}

	unopenSafeCells := game.board.CountUnopenSafeCells()
	allSafeCellsAreOpen := unopenSafeCells == 0

	if allSafeCellsAreOpen {
		return Won
	}

	return Playing
}

func (game *Game) Lives() int {
	return game.lives
}

func (game *Game) tryCascade(cascadingOrigin dimensions.Location) {
	for _, location := range cascadingOrigin.AdjacentLocations() {
		candidateCell := game.board.Retrieve(location)

		if candidateCell == board.Void {
			continue
		}

		if candidateCell.Status(board.Opened, board.Flagged) {
			continue
		}

		if candidateCell.Mined() {
			continue
		}

		// We want to open up until (and including) the first cell with at least one adjacent mine
		// If there are adjacent mines we will stop cascading

		candidateCell.Open()

		if candidateCell.AdjacentMines() == 0 {
			game.tryCascade(candidateCell.Position())
		}
	}
}

func (game *Game) ended() bool {
	if game.Status() == Playing {
		return false
	}

	return true
}
