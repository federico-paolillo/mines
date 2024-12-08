package validators_test

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/server/validators"
	"github.com/go-playground/validator/v10"
)

type difficultyData struct {
	Difficulty string `validate:"isdifficultyenum"`
}

func TestDifficultyValidatorValidatesProperly(t *testing.T) {
	v := validator.New()

	err := v.RegisterValidation(validators.IsDifficultyEnumValidator, validators.IsDifficultyEnum)

	if err != nil {
		t.Errorf(
			"failed to register validation. %v",
			err,
		)
	}

	testCases := []struct {
		data difficultyData
		isOk bool
	}{
		{
			isOk: true,
			data: difficultyData{
				Difficulty: "beginner",
			},
		},
		{
			isOk: true,
			data: difficultyData{
				Difficulty: "intermediate",
			},
		},
		{
			isOk: true,
			data: difficultyData{
				Difficulty: "expert",
			},
		},
		{
			isOk: false,
			data: difficultyData{
				Difficulty: "pippo",
			},
		},
	}

	for _, testCase := range testCases {
		err := v.Struct(testCase.data)

		if (err == nil) != testCase.isOk {
			t.Errorf(
				"unexpected validation result for value '%s'. %v",
				testCase.data.Difficulty,
				err,
			)
		}
	}
}
