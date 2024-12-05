package res

import (
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type CellDto struct {
	State board.Cellstate `json:"state"`
	X     int             `json:"x"`
	Y     int             `json:"y"`
}

func ToCellDto(cell matchmaking.Cell) CellDto {
	return CellDto{
		State: cell.State,
		X:     cell.X,
		Y:     cell.Y,
	}
}
