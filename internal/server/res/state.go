package res

type Gamestate = string

const (
	Playing Gamestate = "playing"
	Won               = "won"
	Lost              = "lost"
)

type GameStateDto struct {
	Id     string      `json:"id"`
	Lives  int         `json:"lives"`
	State  Gamestate   `json:"state"`
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Cells  [][]CellDto `json:"cells"`
}
