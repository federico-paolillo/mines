package res

import (
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type MatchstateDto struct {
	Id        string         `json:"id"`
	Lives     int            `json:"lives"`
	State     game.Gamestate `json:"state"`
	Width     int            `json:"width"`
	Height    int            `json:"height"`
	Cells     []CellDto      `json:"cells"`
	StartTime int64          `json:"startTime"`
}

func ToMatchstateDto(matchstate *matchmaking.Matchstate) MatchstateDto {
	cells := make([]CellDto, 0, matchstate.Width*matchstate.Height)

	for _, row := range matchstate.Cells {
		for _, cell := range row {
			cells = append(
				cells,
				ToCellDto(cell),
			)
		}
	}

	return MatchstateDto{
		Id:        matchstate.Id,
		Lives:     matchstate.Lives,
		State:     matchstate.State,
		Width:     matchstate.Width,
		Height:    matchstate.Height,
		Cells:     cells,
		StartTime: matchstate.StartTime,
	}
}
