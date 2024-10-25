package res

import "github.com/federico-paolillo/mines/pkg/board"

type CellDto struct {
	State board.Cellstate `json:"state"`
	X     int             `json:"x"`
	Y     int             `json:"y"`
}
