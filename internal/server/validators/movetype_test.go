package validators_test

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/server/validators"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/go-playground/validator/v10"
)

type moveTypeData struct {
	MoveType string `validate:"ismovetypeenum"`
}

func TestMoveTypeValidatorValidatesProperly(t *testing.T) {
	v := validator.New()

	err := v.RegisterValidation(validators.IsMoveTypeEnumValidator, validators.IsMoveTypeEnum)

	if err != nil {
		t.Errorf(
			"failed to register validation. %v",
			err,
		)
	}

	testCases := []struct {
		data moveTypeData
		isOk bool
	}{
		{
			isOk: true,
			data: moveTypeData{
				MoveType: string(matchmaking.MoveChord),
			},
		},
		{
			isOk: true,
			data: moveTypeData{
				MoveType: string(matchmaking.MoveFlag),
			},
		},
		{
			isOk: true,
			data: moveTypeData{
				MoveType: string(matchmaking.MoveOpen),
			},
		},
		{
			isOk: false,
			data: moveTypeData{
				MoveType: "pippo",
			},
		},
	}

	for _, testCase := range testCases {
		err := v.Struct(testCase.data)

		if (err == nil) != testCase.isOk {
			t.Errorf(
				"unexpected validation result for value '%s'. %v",
				testCase.data.MoveType,
				err,
			)
		}
	}
}
