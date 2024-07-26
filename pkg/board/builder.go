package board

import "github.com/federico-paolillo/mines/pkg/dimensions"

type Cellkind = string

const (
	mine Cellkind = "mine"
	void          = "void"
	safe          = "safe"
)

type Placements = map[dimensions.Location]Cellkind

type Builder struct {
	size       dimensions.Size
	placements Placements
}

func NewBuilder(boardSize dimensions.Size) *Builder {
	return &Builder{
		size:       boardSize,
		placements: make(map[dimensions.Location]Cellkind, boardSize.Width*boardSize.Height),
	}
}

func (builder *Builder) PlaceSafe(x, y int) {
	builder.placements[dimensions.Location{X: x, Y: y}] = safe
}

func (builder *Builder) PlaceMine(x, y int) {
	builder.placements[dimensions.Location{X: x, Y: y}] = mine
}

func (builder *Builder) PlaceVoid(x, y int) {
	builder.placements[dimensions.Location{X: x, Y: y}] = void
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
