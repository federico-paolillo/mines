package intents

type Position interface {
	XCoord() int
	YCoord() int
}

type Coords struct {
	X, Y int
}

func (c Coords) XCoord() int {
	return c.X
}

func (c Coords) YCoord() int {
	return c.Y
}

type FlagIntent struct {
	Coords
}

type OpenIntent struct {
	Coords
}

type ChordIntent struct {
	Coords
}
