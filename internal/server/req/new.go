package req

import "github.com/federico-paolillo/mines/pkg/game"

type NewGameDto struct {
	Difficulty game.Difficulty `binding:"required,isdifficultyenum" json:"difficulty"`
}
