package game

import "github.com/federico-paolillo/mines/pkg/mines"

type GameStatus = string

const (
	Playing GameStatus = "playing"
	Lost               = "lost"
	Won                = "won"
)

type Game struct {
	Board  *mines.Board
	Lives  int
	Status GameStatus
}

func NewGame(lives int, board *mines.Board) *Game {
	return &Game{
		Board:  board,
		Lives:  lives,
		Status: Playing,
	}
}

func (game *Game) Flag(location mines.Location) {
	cell := game.Board.Retrieve(location)

	if cell == mines.Void {
		return
	}

	if cell.Status == mines.Opened {
		return
	}

	cell.Status = mines.Flagged

	game.checkWinCondition()
}

func (game *Game) Open(location mines.Location) {
	cell := game.Board.Retrieve(location)

	if cell == mines.Void {
		return
	}

	if cell.Status == mines.Opened {
		return
	}

	if cell.Mined {
		game.mineTripped()
		return
	}

	cell.Status = mines.Opened

	if game.Board.AdjacentMines(cell.Position) == 0 {
		game.tryChording(cell.Position)
	}

	game.checkWinCondition()
}

func (game *Game) tryChording(chordingOrigin mines.Location) {
	for _, location := range chordingOrigin.AdjacentLocations() {
		candidateCell := game.Board.Retrieve(location)

		if candidateCell == mines.Void {
			continue
		}

		if candidateCell.Status == mines.Opened {
			continue
		}

		if candidateCell.Status == mines.Flagged {
			continue
		}

		if candidateCell.Mined {
			continue
		}

		if game.Board.AdjacentMines(candidateCell.Position) == 0 {
			candidateCell.Status = mines.Opened
			game.tryChording(candidateCell.Position)
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
	// All mines are flagged
	// OR
	// All safe cells are open
}
