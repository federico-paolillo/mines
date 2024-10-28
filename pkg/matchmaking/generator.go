package matchmaking

import (
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

type BoardGenerator interface {
	Generate(size dimensions.Size, mines int) *board.Board
}
