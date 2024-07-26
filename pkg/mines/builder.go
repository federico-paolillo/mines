package mines

type Cellkind = string

const (
	mine Cellkind = "mine"
	void          = "void"
	safe          = "safe"
)

type Placements = map[Location]Cellkind

type Builder struct {
	size       Size
	placements Placements
}

func NewBuilder(boardSize Size) *Builder {
	return &Builder{
		size:       boardSize,
		placements: make(map[Location]Cellkind, boardSize.Width*boardSize.Height),
	}
}

func (builder *Builder) PlaceSafe(x, y int) {
	builder.placements[Location{X: x, Y: y}] = safe
}

func (builder *Builder) PlaceMine(x, y int) {
	builder.placements[Location{X: x, Y: y}] = mine
}

func (builder *Builder) PlaceVoid(x, y int) {
	builder.placements[Location{X: x, Y: y}] = void
}

func (builder *Builder) Build() *Board {
	board := NewBoard(builder.size)

	for location, cellkind := range builder.placements {
		switch cellkind {
		case mine:
			board.PlaceMine(location)
		case safe:
			board.PlaceCell(location)
		default:
			board.PlaceVoid(location)
		}
	}

	return board
}
