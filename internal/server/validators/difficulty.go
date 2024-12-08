package validators

import (
	"reflect"

	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/go-playground/validator/v10"
)

const IsDifficultyEnumValidator = "isdifficultyenum"

var IsDifficultyEnum validator.Func = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() == reflect.String {
		v := field.String()

		switch v {
		case string(game.BeginnerDifficulty):
			fallthrough
		case string(game.IntermediateDifficulty):
			fallthrough
		case string(game.ExpertDifficulty):
			return true
		default:
			return false
		}
	}

	return false
}
