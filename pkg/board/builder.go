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
			cells[location] = NewMineCell(location, builder.countAdjacentMines(location))
		case safe:
			cells[location] = NewSafeCell(location, builder.countAdjacentMines(location))
		default:
			continue
		}
	}

	board := newBoard(cells)

	return board
}

func (builder *Builder) countAdjacentMines(location dimensions.Location) int {
	adjacentLocations := location.AdjacentLocations()
	minesCount := 0

	for _, adjacentLocation := range adjacentLocations {
		if val, ok := builder.placements[adjacentLocation]; ok {
			if val == mine {
				minesCount++
			}
		}
	}

	return minesCount
}
