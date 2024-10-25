package res

import (
	"github.com/federico-paolillo/mines/pkg/game"
)

type MatchstateDto struct {
	Id     string         `json:"id"`
	Lives  int            `json:"lives"`
	State  game.Gamestate `json:"state"`
	Width  int            `json:"width"`
	Height int            `json:"height"`
	Cells  [][]CellDto    `json:"cells"`
}
