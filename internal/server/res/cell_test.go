package res_test

import (
	"encoding/json"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/board"
)

func TestCellDtoMarshalsCorrectly(t *testing.T) {
	cellDto := &res.CellDto{
		State: board.FlaggedCell,
		X:     12,
		Y:     23,
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
  "y": 23
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
