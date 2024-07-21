package game

import "github.com/federico-paolillo/mines/pkg/mines"

type Difficulty = string

const (
	Beginner     Difficulty = "beginner"
	Intermediate            = "intermediate"
	Expert                  = "expert"
)

type DifficultySettings struct {
	BoardSize     mines.Size
	NumberOfMines int
	Lives         int
}

var beginnerSettings = DifficultySettings{
	BoardSize: mines.Size{
		Width:  9,
		Height: 9,
	},
	NumberOfMines: 10,
	Lives:         2,
}

var intermediateSettings = DifficultySettings{
	BoardSize: mines.Size{
		Width:  16,
		Height: 16,
	},
	NumberOfMines: 40,
	Lives:         1,
}

var expertSettings = DifficultySettings{
	BoardSize: mines.Size{
		Width:  30,
		Height: 16,
	},
	NumberOfMines: 99,
	Lives:         0,
}

func GetDifficultySettings(difficulty Difficulty) DifficultySettings {
	switch difficulty {
	case Beginner:
		return beginnerSettings
	case Intermediate:
		return intermediateSettings
	case Expert:
		return expertSettings
	default:
		return beginnerSettings
	}
}
