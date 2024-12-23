package matchmaking

type Movetype string

const (
	MoveOpen  Movetype = "open"
	MoveFlag  Movetype = "flag"
	MoveChord Movetype = "chord"
)

type Move struct {
	Type Movetype
	X, Y int
}
