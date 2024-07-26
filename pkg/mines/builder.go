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
	cells := make(Cellmap, builder.size.Width*builder.size.Height)

	for location, cellkind := range builder.placements {
		switch cellkind {
		case mine:
			cells[location] = NewMineCell(location)
		case safe:
			cells[location] = NewSafeCell(location)
		default:
			continue
		}
	}

	board := newBoard(cells)

	return board
}
