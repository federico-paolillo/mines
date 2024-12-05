package req_test

import (
	"encoding/json"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestNewGameDtoUnmarshalsCorrectly(t *testing.T) {
	var newGameDto req.NewGameDto

	newGameDtoJson := `{
  "difficulty": "expert"
}`

	err := json.Unmarshal([]byte(newGameDtoJson), &newGameDto)

	if err != nil {
		t.Fatalf(
			"could not unmarshal move dto. %v",
			err,
		)
	}

	expectedNewGameDto := req.NewGameDto{
		Difficulty: game.ExpertDifficulty,
	}

	if newGameDto != expectedNewGameDto {
		t.Fatalf(
			"unmarshal did not work, stuff is missing. wanted\n%+v\ngot\n%+v\n",
			expectedNewGameDto,
			newGameDto,
		)
	}
}
