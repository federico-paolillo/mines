package req_test

import (
	"encoding/json"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

func TestMoveDtoUnmarshalsCorrectly(t *testing.T) {
	var moveDto req.MoveDto

	moveDtoJson := `{
  "type": "chord",
	"x": 123,
	"y": 321
}`

	err := json.Unmarshal([]byte(moveDtoJson), &moveDto)

	if err != nil {
		t.Fatalf(
			"could not unmarshal move dto. %v",
			err,
		)
	}

	expectedMoveDto := req.MoveDto{
		Type: matchmaking.MoveChord,
		X:    123,
		Y:    321,
	}

	if moveDto != expectedMoveDto {
		t.Fatalf(
			"unmarshal did not work, stuff is missing. wanted\n%+v\ngot\n%+v\n",
			expectedMoveDto,
			moveDto,
		)
	}
}
