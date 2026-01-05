package res_test

import (
	"encoding/json"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

func TestCellDtoMarshalsCorrectly(t *testing.T) {
	cellDto := &res.CellDto{
		State:         board.FlaggedCell,
		X:             12,
		Y:             23,
		AdjacentMines: 100,
	}

	b, err := json.MarshalIndent(cellDto, "", "  ")

	if err != nil {
		t.Errorf(
			"failure when marshalling cell dto. %v",
			err,
		)
	}

	expectedString := `{
  "state": "flagged",
  "x": 12,
  "y": 23,
  "adjacentMines": 100
}`

	actualString := string(b)

	if expectedString != actualString {
		t.Fatalf(
			"cell dto was serialized to something unexpected. wanted\n%s\ngot\n%s\n",
			expectedString,
			actualString,
		)
	}
}

func TestCellDtoMapsCorrectlyFromMatchmakingCell(t *testing.T) {
	cell := matchmaking.Cell{
		X:             1,
		Y:             2,
		Mined:         true,
		State:         board.FlaggedCell,
		AdjacentMines: 14,
	}

	cellDto := res.ToCellDto(cell)

	expectedCellDto := res.CellDto{
		X:             1,
		Y:             2,
		State:         board.FlaggedCell,
		AdjacentMines: 14,
	}

	if cellDto != expectedCellDto {
		t.Fatalf(
			"cell dto was mapped incorrectly. wanted\n%v\ngot\n%v\n",
			expectedCellDto,
			cellDto,
		)
	}
}
