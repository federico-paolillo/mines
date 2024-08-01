package board

import (
	"errors"
	"fmt"

	"github.com/federico-paolillo/mines/pkg/dimensions"
)

type Cellkind = string

const (
	Mine Cellkind = "mine"
	void          = "void"
	Safe          = "safe"
)

type Placements = map[dimensions.Location]Cellkind

type Builder struct {
	size       dimensions.Size
	placements Placements
}

var ErrOutOfBounds = errors.New("placement is out of bounds")

func NewBuilder(boardSize dimensions.Size) *Builder {
	return &Builder{
		size:       boardSize,
		placements: make(map[dimensions.Location]Cellkind, boardSize.Width*boardSize.Height),
	}
}

func (builder *Builder) PlaceSafe(x, y int) error {
	location := dimensions.Location{X: x, Y: y}

	if !builder.size.Contains(location) {
		return fmt.Errorf(
			"%v is out of bounds for size %v. %w",
			location,
			builder.size,
			ErrOutOfBounds,
		)
	}

	builder.placements[location] = Safe

	return nil
}

func (builder *Builder) PlaceMine(x, y int) error {
	location := dimensions.Location{X: x, Y: y}

	if !builder.size.Contains(location) {
		return fmt.Errorf(
			"%v is out of bounds for size %v. %w",
			location,
			builder.size,
			ErrOutOfBounds,
		)
	}

	builder.placements[location] = Mine

	return nil
}

func (builder *Builder) PlaceVoid(x, y int) error {
	location := dimensions.Location{X: x, Y: y}

	if !builder.size.Contains(location) {
		return fmt.Errorf(
			"%v is out of bounds for size %v. %w",
			location,
			builder.size,
			ErrOutOfBounds,
		)
	}

	builder.placements[location] = void

	return nil
}

func (builder *Builder) Build() *Board {
	cells := make(Cellmap, builder.size.Width*builder.size.Height)

	for location, cellkind := range builder.placements {
		switch cellkind {
		case Mine:
			cells[location] = NewMineCell(location, builder.countAdjacentMines(location))
		case Safe:
			cells[location] = NewSafeCell(location, builder.countAdjacentMines(location))
		default:
			continue
		}
	}

	board := newBoard(builder.size, cells)

	return board
}

func (builder *Builder) IsSafe(x, y int) bool {
	return builder.getAt(x, y) == Safe
}

func (builder *Builder) IsMine(x, y int) bool {
	return builder.getAt(x, y) == Mine
}

func (builder *Builder) getAt(x, y int) Cellkind {
	location := dimensions.Location{X: x, Y: y}

	if cellkind, ok := builder.placements[location]; ok {
		return cellkind
	}

	return void
}

func (builder *Builder) countAdjacentMines(location dimensions.Location) int {
	adjacentLocations := location.AdjacentLocations()
	minesCount := 0

	for _, adjacentLocation := range adjacentLocations {
		if val, ok := builder.placements[adjacentLocation]; ok {
			if val == Mine {
				minesCount++
			}
		}
	}

	return minesCount
}
