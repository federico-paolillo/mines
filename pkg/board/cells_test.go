package board_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

func TestNewSafeCellIsNotMinedAndPositionedProperly(t *testing.T) {
	emptyCellLocation := dimensions.Location{X: 1, Y: 2}
	emptyCell := board.NewSafeCell(emptyCellLocation)

	if emptyCell.Position() != emptyCellLocation {
		t.Errorf(
			"expected cell to be at %v. instead is at %v",
			emptyCellLocation,
			emptyCell.Position(),
		)
	}

	if emptyCell.Mined() {
		t.Error("expected cell to be without mines")
	}

	if emptyCell.Status(board.Closed) != true {
		t.Error("expected cell to be closed")
	}
}

func TestNewMineCellIsMinedAndPositionedProperly(t *testing.T) {
	minedCellLocation := dimensions.Location{X: 1, Y: 2}
	minedCell := board.NewMineCell(minedCellLocation)

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

	if minedCell.Status(board.Closed) != true {
		t.Error("expected cell to be closed")
	}
}
