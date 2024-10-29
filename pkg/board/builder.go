package board

import (
	"errors"
	"fmt"

	"github.com/federico-paolillo/mines/pkg/dimensions"
)

type Cellkind = string

const (
	MineKind Cellkind = "mine"
	voidKind          = "void"
	SafeKind          = "safe"
)

type metaCell struct {
	kind  Cellkind
	state Cellstate
}

var void = metaCell{
	voidKind,
	ClosedCell,
}

type Placements = map[dimensions.Location]metaCell

type Builder struct {
	size       dimensions.Size
	placements Placements
}

var ErrOutOfBounds = errors.New("placement is out of bounds")

func NewBuilder(boardSize dimensions.Size) *Builder {
	return &Builder{
		size:       boardSize,
		placements: make(map[dimensions.Location]metaCell, boardSize.Width*boardSize.Height),
	}
}

func (builder *Builder) PlaceSafe(x, y int) error {
	return builder.place(
		x,
		y,
		SafeKind,
		ClosedCell,
	)
}

func (builder *Builder) PlaceMine(x, y int) error {
	return builder.place(
		x,
		y,
		MineKind,
		ClosedCell,
	)
}

func (builder *Builder) PlaceVoid(x, y int) error {
	return builder.place(
		x,
		y,
		voidKind,
		UnfathomableCell,
	)
}

func (builder *Builder) MarkOpen(x, y int) error {
	return builder.mark(
		x,
		y,
		OpenCell,
	)
}

func (builder *Builder) MarkClose(x, y int) error {
	return builder.mark(
		x,
		y,
		ClosedCell,
	)
}

func (builder *Builder) MarkFlag(x, y int) error {
	return builder.mark(
		x,
		y,
		FlaggedCell,
	)
}

func (builder *Builder) mark(
	x, y int,
	newState Cellstate,
) error {
	location := dimensions.Location{X: x, Y: y}

	if !builder.size.Contains(location) {
		return fmt.Errorf(
			"builder: '%v' is out of bounds for size '%v'. %w",
			location,
			builder.size,
			ErrOutOfBounds,
		)
	}

	if cell, ok := builder.placements[location]; ok {
		cell.state = newState
	}

	return nil
}

func (builder *Builder) place(
	x, y int,
	kind Cellkind,
	state Cellstate,
) error {
	location := dimensions.Location{X: x, Y: y}

	if !builder.size.Contains(location) {
		return fmt.Errorf(
			"builder: '%v' is out of bounds for size '%v'. %w",
			location,
			builder.size,
			ErrOutOfBounds,
		)
	}

	builder.placements[location] = metaCell{
		kind:  kind,
		state: state,
	}

	return nil
}

func (builder *Builder) Build() *Board {
	cells := make(Cellmap, builder.size.Width*builder.size.Height)

	for location, cellmeta := range builder.placements {
		switch cellmeta {
		case MineKind:
			cells[location] = NewMineCell(location, builder.countAdjacentMines(location))
		case SafeKind:
			cells[location] = NewSafeCell(location, builder.countAdjacentMines(location))
		default:
			continue
		}
	}

	board := newBoard(builder.size, cells)

	return board
}

func (builder *Builder) IsSafe(x, y int) bool {
	return builder.getAt(x, y).kind == SafeKind
}

func (builder *Builder) IsMine(x, y int) bool {
	return builder.getAt(x, y).kind == MineKind
}

func (builder *Builder) getAt(x, y int) metaCell {
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
		if cell, ok := builder.placements[adjacentLocation]; ok {
			if cell.kind == MineKind {
				minesCount++
			}
		}
	}

	return minesCount
}
