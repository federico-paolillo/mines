package res_test

import (
	"encoding/json"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestStateDtoMarshalsCorrectly(t *testing.T) {
	stateDto := res.MatchstateDto{
		Id:     "abc",
		Lives:  13,
		State:  game.Lost,
		Width:  123,
		Height: 321,
		Cells: [][]res.CellDto{
			{
				res.CellDto{
					State: board.Flagged,
					X:     1,
					Y:     1,
				},
				res.CellDto{
					State: board.Flagged,
					X:     2,
					Y:     1,
				},
				res.CellDto{
					State: board.Flagged,
					X:     3,
					Y:     1,
				},
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
    [
      {
        "state": "flagged",
        "x": 1,
        "y": 1
      },
      {
        "state": "flagged",
        "x": 2,
        "y": 1
      },
      {
        "state": "flagged",
        "x": 3,
        "y": 1
      }
    ]
  ]
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
