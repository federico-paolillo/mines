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
	board  *board.Board
	lives  int
	status GameStatus
}

func NewGame(lives int, board *board.Board) *Game {
	return &Game{
		board:  board,
		lives:  lives,
		status: Playing,
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
		game.mineTripped()
		return
	}

	if cell.AdjacentMines() == 0 {
		game.tryCascade(cell.Position())
	}

	game.checkWinCondition()
}

func (game *Game) Status() GameStatus {
	return game.status
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

func (game *Game) mineTripped() {
	if game.lives == 0 {
		game.status = Lost
	} else {
		game.lives--
	}
}

func (game *Game) checkWinCondition() {
	unopenSafeCells := game.board.CountUnopenSafeCells()

	allSafeCellsAreOpen := unopenSafeCells == 0

	if allSafeCellsAreOpen {
		game.status = Won
		return
	}
}

func (game *Game) ended() bool {
	if game.status == Won {
		return true
	}

	if game.status == Lost {
		return true
	}

	return false
}
