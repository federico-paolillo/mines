package res_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

func TestStateDtoMarshalsCorrectly(t *testing.T) {
	stateDto := res.MatchstateDto{
		Id:        "abc",
		Lives:     13,
		State:     game.LostGame,
		Width:     123,
		Height:    321,
		StartTime: 1700000000,
		Cells: []res.CellDto{
			{
				State:         board.FlaggedCell,
				X:             1,
				Y:             1,
				AdjacentMines: 1,
			},
			{
				State:         board.FlaggedCell,
				X:             2,
				Y:             1,
				AdjacentMines: 2,
			},
			{
				State:         board.FlaggedCell,
				X:             3,
				Y:             1,
				AdjacentMines: 3,
			},
		},
	}

	b, err := json.MarshalIndent(stateDto, "", "  ")

	if err != nil {
		t.Errorf(
			"failure when marshalling state dto. %v",
			err,
		)
	}

	expectedString := `{
  "id": "abc",
  "lives": 13,
  "state": "lost",
  "width": 123,
  "height": 321,
  "cells": [
    {
      "state": "flagged",
      "x": 1,
      "y": 1,
      "adjacentMines": 1
    },
    {
      "state": "flagged",
      "x": 2,
      "y": 1,
      "adjacentMines": 2
    },
    {
      "state": "flagged",
      "x": 3,
      "y": 1,
      "adjacentMines": 3
    }
  ],
  "startTime": 1700000000
}`

	actualString := string(b)

	if expectedString != actualString {
		t.Fatalf(
			"state dto was serialized to something unexpected. wanted\n%s\ngot\n%s\n",
			expectedString,
			actualString,
		)
	}
}

func TestStateDtoMapsCorrectlyFromMatchmakingMatchstate(t *testing.T) {
	matchstate := &matchmaking.Matchstate{
		Id:        "abc",
		Version:   1234,
		Lives:     13,
		State:     game.LostGame,
		Width:     3,
		Height:    1,
		StartTime: 1700000000,
		Cells: [][]matchmaking.Cell{
			{
				matchmaking.Cell{
					X:             1,
					Y:             1,
					Mined:         false,
					State:         board.FlaggedCell,
					AdjacentMines: 11,
				},
				matchmaking.Cell{
					X:             2,
					Y:             1,
					Mined:         false,
					State:         board.FlaggedCell,
					AdjacentMines: 12,
				},
				matchmaking.Cell{
					X:             3,
					Y:             1,
					Mined:         false,
					State:         board.FlaggedCell,
					AdjacentMines: 13,
				},
			},
		},
	}

	matchstateDto := res.ToMatchstateDto(matchstate)

	expectedStateDto := res.MatchstateDto{
		Id:        "abc",
		Lives:     13,
		State:     game.LostGame,
		Width:     3,
		Height:    1,
		StartTime: 1700000000,
		Cells: []res.CellDto{
			{
				State:         board.FlaggedCell,
				X:             1,
				Y:             1,
				AdjacentMines: 11,
			},
			{
				State:         board.FlaggedCell,
				X:             2,
				Y:             1,
				AdjacentMines: 12,
			},
			{
				State:         board.FlaggedCell,
				X:             3,
				Y:             1,
				AdjacentMines: 13,
			},
		},
	}

	if !reflect.DeepEqual(matchstateDto, expectedStateDto) {
		t.Fatalf(
			"matchstate dto was mapped incorrectly. wanted\n%v\ngot\n%v\n",
			matchstateDto,
			expectedStateDto,
		)
	}
}
