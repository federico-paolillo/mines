package matchmaking_test

import (
	"reflect"
	"testing"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

func TestExportsCellsProperly(t *testing.T) {
	s := dimensions.Size{
		Width:  4,
		Height: 4,
	}

	bb := board.NewBuilder(s)

	for x := range s.Width {
		for y := range s.Height {
			_ = bb.PlaceSafe(x+1, y+1)
		}
	}

	bb.PlaceMine(1, 1)

	b := bb.Build()

	expectation := [][]matchmaking.Cell{
		{
			{
				X:             1,
				Y:             1,
				State:         board.ClosedCell,
				Mined:         true,
				AdjacentMines: 0,
			},
			{
				X:             2,
				Y:             1,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 1,
			},
			{
				X:             3,
				Y:             1,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
			{
				X:             4,
				Y:             1,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
		},
		{
			{
				X:             1,
				Y:             2,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 1,
			},
			{
				X:             2,
				Y:             2,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 1,
			},
			{
				X:             3,
				Y:             2,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
			{
				X:             4,
				Y:             2,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
		},
		{
			{
				X:             1,
				Y:             3,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
			{
				X:             2,
				Y:             3,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
			{
				X:             3,
				Y:             3,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
			{
				X:             4,
				Y:             3,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
		},
		{
			{
				X:             1,
				Y:             4,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
			{
				X:             2,
				Y:             4,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
			{
				X:             3,
				Y:             4,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
			{
				X:             4,
				Y:             4,
				State:         board.ClosedCell,
				Mined:         false,
				AdjacentMines: 0,
			},
		},
	}

	cells := matchmaking.ExportCells(b)

	sameCells := reflect.DeepEqual(cells, expectation)

	if !sameCells {
		t.Fatalf(
			"cells differ. wanted\n%v\ngot\n%v\n",
			expectation,
			cells,
		)
	}
}
