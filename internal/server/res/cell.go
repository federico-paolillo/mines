package res

type Cellstate = string

const (
	Open    Cellstate = "open"
	Closed            = "closed"
	Flagged           = "flagged"
)

type CellDto struct {
	State Cellstate `json:"state"`
	X     int       `json:"x"`
	Y     int       `json:"y"`
}
