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

func (game *Game) Flag(location dimensions.Location) {
	cell := game.Board.Retrieve(location)

	if cell == board.Void {
		return
	}

	cell.Flag()

	game.checkWinCondition()
}

func (game *Game) Open(location dimensions.Location) {
	cell := game.Board.Retrieve(location)

	if cell == board.Void {
		return
	}

	cell.Open()

	if cell.Mined() {
		game.mineTripped()
		return
	}

	if game.Board.AdjacentMines(cell.Position()) == 0 {
		game.tryChording(cell.Position())
	}

	game.checkWinCondition()
}

func (game *Game) tryChording(chordingOrigin dimensions.Location) {
	for _, location := range chordingOrigin.AdjacentLocations() {
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

		if game.Board.AdjacentMines(candidateCell.Position()) == 0 {
			candidateCell.Open()
			game.tryChording(candidateCell.Position())
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
