package req

import "github.com/federico-paolillo/mines/pkg/matchmaking"

type MoveDto struct {
	Type matchmaking.Movetype `binding:"required,ismovetypeenum" json:"type"`
	X    int                  `binding:"required"                json:"x"`
	Y    int                  `binding:"required"                json:"y"`
}
