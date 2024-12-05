package res

import (
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type MatchstateDto struct {
	Id     string         `json:"id"`
	Lives  int            `json:"lives"`
	State  game.Gamestate `json:"state"`
	Width  int            `json:"width"`
	Height int            `json:"height"`
	Cells  [][]CellDto    `json:"cells"`
}

func ToMatchstateDto(matchstate *matchmaking.Matchstate) MatchstateDto {
	rows := make([][]CellDto, 0, len(matchstate.Cells))

	for _, row := range matchstate.Cells {
		cols := make([]CellDto, 0, len(row))

		for _, cell := range row {
			cols = append(
				cols,
				ToCellDto(cell),
			)
		}

		rows = append(rows, cols)
	}

	return MatchstateDto{
		Id:     matchstate.Id,
		Lives:  matchstate.Lives,
		State:  matchstate.State,
		Width:  matchstate.Width,
		Height: matchstate.Height,
		Cells:  rows,
	}
}
