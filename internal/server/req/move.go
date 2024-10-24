package req

type Movetype = string

const (
	Open  Movetype = "open"
	Flag           = "flag"
	Chord          = "chord"
)

type MoveDto struct {
	Type Movetype `json:"type"`
	X    int      `json:"x"`
	Y    int      `json:"y"`
}
