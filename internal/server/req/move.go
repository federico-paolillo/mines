package req

type Movetype = string

const (
	Open  Movetype = "open"
	Flag  Movetype = "flag"
	Chord Movetype = "chord"
)

type MoveDto struct {
	Type Movetype `json:"type"`
	X    int      `json:"x"`
	Y    int      `json:"y"`
}
