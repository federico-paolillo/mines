package mines_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/mines"
)

func TestNewEmptyCellIsNotMinedAndPositionedProperly(t *testing.T) {
	emptyCellLocation := mines.Location{1, 2}
	emptyCell := mines.NewEmptyCell(emptyCellLocation)

	if emptyCell.Position != emptyCellLocation {
		t.Errorf("expected cell to be at %v. instead is at %v", emptyCellLocation, emptyCell.Position)
	}

	if emptyCell.Mined != false {
		t.Error("expected cell to be without mines")
	}

	if emptyCell.Status != mines.Closed {
		t.Error("expected cell to be closed")
	}
}

func TestNewMineCellIsMinedAndPositionedProperly(t *testing.T) {
	minedCellLocation := mines.Location{1, 2}
	minedCell := mines.NewMinedCell(minedCellLocation)

	if minedCell.Position != minedCellLocation {
		t.Errorf("expected cell to be at %v. instead is at %v", minedCellLocation, minedCell.Position)
	}

	if minedCell.Mined != true {
		t.Error("expected cell to be mined")
	}

	if minedCell.Status != mines.Closed {
		t.Error("expected cell to be closed")
	}
}
