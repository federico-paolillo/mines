package memcached_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/federico-paolillo/mines/internal/storage/memcached"
	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/stretchr/testify/assert"
)

func TestMatchstateToJson(t *testing.T) {
	matchstate := &matchmaking.Matchstate{
		Id:      "test-id",
		Version: 123,
		Lives:   3,
		State:   game.PlayingGame,
		Width:   2,
		Height:  2,
		Cells: matchmaking.Cells{
			{
				{X: 0, Y: 0, Mined: false, State: board.OpenCell},
				{X: 1, Y: 0, Mined: true, State: board.ClosedCell},
			},
			{
				{X: 0, Y: 1, Mined: false, State: board.FlaggedCell},
				{X: 1, Y: 1, Mined: false, State: board.OpenCell},
			},
		},
		StartTime: time.Now().Unix(),
	}

	jsonState := memcached.MatchstateToJSON(matchstate)

	assert.NotNil(t, jsonState)
	assert.Equal(t, matchstate.Id, jsonState.Id)
	assert.Equal(t, matchstate.Lives, jsonState.Lives)
	assert.Equal(t, string(matchstate.State), jsonState.State)
	assert.Equal(t, matchstate.Width, jsonState.Width)
	assert.Equal(t, matchstate.Height, jsonState.Height)
	assert.Equal(t, matchstate.StartTime, jsonState.StartTime)

	expectedCells := []memcached.CellJSON{
		{X: 0, Y: 0, Mined: false, State: string(board.OpenCell)},
		{X: 1, Y: 0, Mined: true, State: string(board.ClosedCell)},
		{X: 0, Y: 1, Mined: false, State: string(board.FlaggedCell)},
		{X: 1, Y: 1, Mined: false, State: string(board.OpenCell)},
	}

	assert.ElementsMatch(t, expectedCells, jsonState.Cells)

	// Test JSON marshaling
	_, err := json.Marshal(jsonState)
	assert.NoError(t, err)
}

func TestJsonToMatchstate(t *testing.T) {
	jsonState := &memcached.MatchstateJSON{
		Id:     "test-id",
		Lives:  3,
		State:  string(game.PlayingGame),
		Width:  2,
		Height: 2,
		Cells: []memcached.CellJSON{
			{X: 0, Y: 0, Mined: false, State: string(board.OpenCell)},
			{X: 1, Y: 0, Mined: true, State: string(board.ClosedCell)},
			{X: 0, Y: 1, Mined: false, State: string(board.FlaggedCell)},
			{X: 1, Y: 1, Mined: false, State: string(board.OpenCell)},
		},
		StartTime: time.Now().Unix(),
	}

	matchstate := memcached.JSONToMatchstate(jsonState)

	assert.NotNil(t, matchstate)
	assert.Equal(t, jsonState.Id, matchstate.Id)
	assert.Equal(t, jsonState.Lives, matchstate.Lives)
	assert.Equal(t, game.Gamestate(jsonState.State), matchstate.State)
	assert.Equal(t, jsonState.Width, matchstate.Width)
	assert.Equal(t, jsonState.Height, matchstate.Height)
	assert.Equal(t, jsonState.StartTime, matchstate.StartTime)

	// Reconstruct the expected Cells array in matchmaking.Cells format
	expectedCells := matchmaking.Cells{
		{
			{X: 0, Y: 0, Mined: false, State: board.OpenCell},
			{X: 1, Y: 0, Mined: true, State: board.ClosedCell},
		},
		{
			{X: 0, Y: 1, Mined: false, State: board.FlaggedCell},
			{X: 1, Y: 1, Mined: false, State: board.OpenCell},
		},
	}

	for y := 0; y < jsonState.Height; y++ {
		for x := 0; x < jsonState.Width; x++ {
			assert.Equal(t, expectedCells[y][x].X, matchstate.Cells[y][x].X)
			assert.Equal(t, expectedCells[y][x].Y, matchstate.Cells[y][x].Y)
			assert.Equal(t, expectedCells[y][x].Mined, matchstate.Cells[y][x].Mined)
			assert.Equal(t, expectedCells[y][x].State, matchstate.Cells[y][x].State)
		}
	}

	// Test JSON unmarshaling
	marshaledJSON, err := json.Marshal(jsonState)
	assert.NoError(t, err)

	var unmarshaled memcached.MatchstateJSON
	err = json.Unmarshal(marshaledJSON, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, jsonState.Id, unmarshaled.Id)
	assert.Equal(t, jsonState.Lives, unmarshaled.Lives)
	assert.Equal(t, jsonState.State, unmarshaled.State)
	assert.Equal(t, jsonState.Width, unmarshaled.Width)
	assert.Equal(t, jsonState.Height, unmarshaled.Height)
	assert.Equal(t, jsonState.StartTime, unmarshaled.StartTime)
	assert.ElementsMatch(t, jsonState.Cells, unmarshaled.Cells)
}

func TestJsonToMatchstateNoCells(t *testing.T) {
	match := testutils.NewMatchState(t, "new-match", 0)

	matchToJSON := memcached.MatchstateToJSON(match)

	_ = memcached.JSONToMatchstate(matchToJSON)
}
