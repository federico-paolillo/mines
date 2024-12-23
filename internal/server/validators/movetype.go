package validators

import (
	"reflect"

	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/go-playground/validator/v10"
)

const IsMoveTypeEnumValidator = "ismovetypeenum"

var IsMoveTypeEnum validator.Func = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() == reflect.String {
		v := field.String()

		switch v {
		case string(matchmaking.MoveOpen):
			fallthrough
		case string(matchmaking.MoveFlag):
			fallthrough
		case string(matchmaking.MoveChord):
			return true
		default:
			return false
		}
	}

	return false
}
