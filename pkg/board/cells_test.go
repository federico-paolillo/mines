package board_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

func TestNewSafeCellIsNotMinedAndPositionedProperly(t *testing.T) {
	safeCellLocation := dimensions.Location{X: 1, Y: 2}
	safeCell := board.NewSafeCell(safeCellLocation, 11)

	if safeCell.Position() != safeCellLocation {
		t.Errorf(
			"expected cell to be at %v. instead is at %v",
			safeCellLocation,
			safeCell.Position(),
		)
	}

	if safeCell.Mined() {
		t.Error("expected cell to be without mines")
	}

	if safeCell.HasStatus(board.ClosedCell) != true {
		t.Error("expected cell to be closed")
	}

	if safeCell.AdjacentMines() != 11 {
		t.Errorf("expected cell to have %d adjacent mines. instead it has %d", 11, safeCell.AdjacentMines())
	}
}

func TestNewMineCellIsMinedAndPositionedProperly(t *testing.T) {
	minedCellLocation := dimensions.Location{X: 1, Y: 2}
	minedCell := board.NewMineCell(minedCellLocation, 14)

	if minedCell.Position() != minedCellLocation {
		t.Errorf(
			"expected cell to be at %v. instead is at %v",
			minedCellLocation,
			minedCell.Position(),
		)
	}

	if minedCell.Safe() {
		t.Error("expected cell to be mined")
	}

	if minedCell.HasStatus(board.ClosedCell) != true {
		t.Error("expected cell to be closed")
	}

	if minedCell.AdjacentMines() != 14 {
		t.Errorf("expected cell to have %d adjacent mines. instead it has %d", 14, minedCell.AdjacentMines())
	}
}
