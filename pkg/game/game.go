package game

import (
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

type GameStatus = string

const (
	Playing GameStatus = "playing"
	Lost               = "lost"
	Won                = "won"
)

type Game struct {
	Board  *board.Board
	Lives  int
	Status GameStatus
}

func NewGame(lives int, board *board.Board) *Game {
	return &Game{
		Board:  board,
		Lives:  lives,
		Status: Playing,
	}
}

func (game *Game) Flag(x, y int) {
	cell := game.Board.Retrieve(dimensions.Location{X: x, Y: y})

	if cell == board.Void {
		return
	}

	cell.Flag()

	game.checkWinCondition()
}

func (game *Game) Open(x, y int) {
	cell := game.Board.Retrieve(dimensions.Location{X: x, Y: y})

	if cell == board.Void {
		return
	}

	cell.Open()

	if cell.Mined() {
		game.mineTripped()
		return
	}

	if cell.AdjacentMines() == 0 {
		game.tryCascade(cell.Position())
	}

	game.checkWinCondition()
}

func (game *Game) tryCascade(cascadingOrigin dimensions.Location) {
	for _, location := range cascadingOrigin.AdjacentLocations() {
		candidateCell := game.Board.Retrieve(location)

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

func (game *Game) mineTripped() {
	if game.Lives == 0 {
		game.Status = Lost
	} else {
		game.Lives--
	}
}

func (game *Game) checkWinCondition() {
	unflaggedMines := game.Board.CountUnflaggedMines()
	unopenSafeCells := game.Board.CountUnopenSafeCells()

	noMissingMinesToFlag := unflaggedMines == 0
	noMissingUnopenCells := unopenSafeCells == 0

	if noMissingMinesToFlag {
		game.Status = Won
		return
	}

	if noMissingUnopenCells {
		game.Status = Won
		return
	}
}

func (game *Game) checkLoseCondition() {

}
