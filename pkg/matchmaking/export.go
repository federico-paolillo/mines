package matchmaking

import (
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

func ExportCells(b *board.Board) Cells {
	bSize := b.Size()

	rows := make([][]Cell, 0, bSize.Height)

	for y := range bSize.Height {
		cols := make([]Cell, 0, bSize.Width)

		for x := range bSize.Width {
			location := dimensions.Location{X: x + 1, Y: y + 1}

			bCell := b.Retrieve(location)

			position := bCell.Position()

			cell := Cell{
				X:             position.X,
				Y:             position.Y,
				State:         bCell.Status(),
				Mined:         bCell.Mined(),
				AdjacentMines: bCell.AdjacentMines(),
			}

			cols = append(cols, cell)
		}

		rows = append(rows, cols)
	}

	return rows
}
