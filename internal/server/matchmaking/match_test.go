package matchmaking_test

import (
	"reflect"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/matchmaking"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestMatchStateReflectsGameAndBoardSituation(t *testing.T) {
	s := dimensions.Size{
		Width:  2,
		Height: 2,
	}

	bb := board.NewBuilder(s)

	for x := range s.Width {
		for y := range s.Height {
			_ = bb.PlaceSafe(x+1, y+1)
		}
	}

	b := bb.Build()

	g := game.NewGame(2, b)

	m := matchmaking.NewMatch("abc", b, g)

	expectation := matchmaking.Matchstate{
		Id:     "abc",
		Lives:  2,
		State:  game.PlayingGame,
		Width:  2,
		Height: 2,
		Cells: [][]matchmaking.Cell{
			{
				{
					X:     1,
					Y:     1,
					State: board.ClosedCell,
				},
				{
					X:     2,
					Y:     1,
					State: board.ClosedCell,
				},
			},
			{
				{
					X:     1,
					Y:     2,
					State: board.ClosedCell,
				},
				{
					X:     2,
					Y:     2,
					State: board.ClosedCell,
				},
			},
		},
	}

	state := m.Status()

	areEqual := reflect.DeepEqual(state, expectation)

	if !areEqual {
		t.Fatalf(
			"state differ. wanted\n%v\ngot\n%v\n",
			expectation,
			state,
		)
	}
}
