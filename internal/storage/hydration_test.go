package storage_test

import (
	"reflect"
	"testing"

	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/stretchr/testify/assert"
)

func TestHydrationRestoresMatchProperly(t *testing.T) {
	bb := board.NewBuilder(
		dimensions.Size{
			Width:  4,
			Height: 1,
		},
	)

	bb.PlaceSafe(1, 1)
	bb.PlaceMine(2, 1)
	bb.PlaceMine(3, 1)
	bb.PlaceMine(4, 1)

	bb.MarkOpen(1, 1)
	bb.MarkFlagged(4, 1)

	b := bb.Build()

	g := game.NewGame(12, b)

	m := matchmaking.NewMatch(
		"abc",
		123,
		456,
		b,
		g,
	)

	state := m.Status()

	mReborn := storage.HydrateMatch(state)

	rebornState := mReborn.Status()

	stateMatches := reflect.DeepEqual(state, rebornState)

	if !stateMatches {
		t.Fatalf(
			"hydration produced a different state. wanted\n%+v\ngot\n%+v\n",
			state,
			rebornState,
		)
	}
}

func TestHydrationIgnoresAdjacentMines(t *testing.T) {
	state := &matchmaking.Matchstate{
		Id:        "abc",
		Version:   123,
		StartTime: 456,
		Lives:     12,
		State:     game.PlayingGame,
		Width:     2,
		Height:    2,
		Cells: [][]matchmaking.Cell{
			{
				{
					X:             1,
					Y:             1,
					State:         board.ClosedCell,
					Mined:         false,
					AdjacentMines: 10000,
				},
				{
					X:             2,
					Y:             1,
					State:         board.ClosedCell,
					Mined:         false,
					AdjacentMines: 20000,
				},
			},
			{
				{
					X:             1,
					Y:             2,
					State:         board.ClosedCell,
					Mined:         false,
					AdjacentMines: 30000,
				},
				{
					X:             2,
					Y:             2,
					State:         board.ClosedCell,
					Mined:         false,
					AdjacentMines: 40000,
				},
			},
		},
	}

	mReborn := storage.HydrateMatch(state)

	rebornState := mReborn.Status()

	assert.Zero(t, rebornState.Cells[0][0].AdjacentMines)
	assert.Zero(t, rebornState.Cells[0][1].AdjacentMines)
	assert.Zero(t, rebornState.Cells[1][0].AdjacentMines)
	assert.Zero(t, rebornState.Cells[1][1].AdjacentMines)
}
