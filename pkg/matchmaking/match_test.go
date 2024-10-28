package matchmaking_test

import (
	"reflect"
	"testing"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
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

	m := matchmaking.NewMatch(
		"abc",
		123,
		b,
		g,
	)

	expectation := &matchmaking.Matchstate{
		Id:      "abc",
		Version: 123,
		Lives:   2,
		State:   game.PlayingGame,
		Width:   2,
		Height:  2,
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

func TestMatchAppliesMovesProperly(t *testing.T) {
	s := dimensions.Size{
		Width:  2,
		Height: 2,
	}

	bb := board.NewBuilder(s)

	bb.PlaceSafe(1, 1)
	bb.PlaceSafe(2, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceMine(2, 2)

	b := bb.Build()

	g := game.NewGame(2, b)

	moves := []matchmaking.Move{
		{
			Type: matchmaking.MoveOpen,
			X:    1,
			Y:    1,
		},
		{
			Type: matchmaking.MoveOpen,
			X:    2,
			Y:    1,
		},
		{
			Type: matchmaking.MoveOpen,
			X:    1,
			Y:    2,
		},
	}

	m := matchmaking.NewMatch(
		"abc",
		123,
		b,
		g,
	)

	for i, move := range moves {
		err := m.Apply(move)

		if err != nil {
			t.Errorf(
				"could not apply move '%d'. %v",
				i,
				err,
			)
		}
	}

	matchStatus := m.Status()

	expectedGamestate := game.WonGame
	gamestate := matchStatus.State

	if matchStatus.State != game.WonGame {
		t.Fatalf(
			"match is not in expected state. got '%s' wanted '%s'",
			gamestate,
			expectedGamestate,
		)
	}
}
