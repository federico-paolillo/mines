package game

import "github.com/federico-paolillo/mines/pkg/dimensions"

type Difficulty string

const (
	BeginnerDifficulty     Difficulty = "beginner"
	IntermediateDifficulty Difficulty = "intermediate"
	ExpertDifficulty       Difficulty = "expert"
)

type DifficultySettings struct {
	BoardSize     dimensions.Size
	NumberOfMines int
	Lives         int
}

var beginnerSettings = DifficultySettings{
	BoardSize: dimensions.Size{
		Width:  9,
		Height: 9,
	},
	NumberOfMines: 10,
	Lives:         2,
}

var intermediateSettings = DifficultySettings{
	BoardSize: dimensions.Size{
		Width:  16,
		Height: 16,
	},
	NumberOfMines: 40,
	Lives:         1,
}

var expertSettings = DifficultySettings{
	BoardSize: dimensions.Size{
		Width:  30,
		Height: 16,
	},
	NumberOfMines: 99,
	Lives:         0,
}

func GetDifficultySettings(difficulty Difficulty) DifficultySettings {
	switch difficulty {
	case BeginnerDifficulty:
		return beginnerSettings
	case IntermediateDifficulty:
		return intermediateSettings
	case ExpertDifficulty:
		return expertSettings
	default:
		return beginnerSettings
	}
}
