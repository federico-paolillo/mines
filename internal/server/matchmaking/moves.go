package matchmaking

type Movetype = string

const (
	MoveOpen  Movetype = "open"
	MoveFlag           = "flag"
	MoveChord          = "chord"
)

type Move struct {
	Type Movetype
	X, Y int
}
