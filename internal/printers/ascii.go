package printers

import (
	"errors"
	"fmt"
	"maps"
	"strings"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

type CellRenderKind int

const (
	OpenMine CellRenderKind = iota
	Closed
	Flagged
	OpenSafe
	OpenSafe1
	OpenSafe2
	OpenSafe3
	OpenSafe4
	OpenSafe5
	OpenSafe6
	OpenSafe7
	OpenSafe8
)

type SymbolsMap map[CellRenderKind]string

type AsciiPrinter struct {
	symbolsMap SymbolsMap
}

var defaultSymbolsMap = SymbolsMap{
	OpenMine:  "x",
	Closed:    "o",
	Flagged:   "f",
	OpenSafe:  "0",
	OpenSafe1: "1",
	OpenSafe2: "2",
	OpenSafe3: "3",
	OpenSafe4: "4",
	OpenSafe5: "5",
	OpenSafe6: "6",
	OpenSafe7: "7",
	OpenSafe8: "8",
}

var UnmappedSymbolErr = errors.New("symbol mapping missing")

const eol = '\n'

func NewAsciiPrinter() *AsciiPrinter {
	return &AsciiPrinter{
		symbolsMap: maps.Clone(defaultSymbolsMap),
	}
}

func (asciiPrinter *AsciiPrinter) Render(board *board.Board) (string, error) {
	var sb strings.Builder

	size := board.Size()

	for y := range size.Height {
		for x := range size.Width {
			cell := board.Retrieve(dimensions.Location{X: x + 1, Y: y + 1}) // Boards start at 1,1
			symbol := asSymbol(cell)

			if stringRepr, ok := asciiPrinter.symbolsMap[symbol]; ok {
				_, err := sb.WriteString(stringRepr)

				if err != nil {
					return "", fmt.Errorf(
						"asciiprinter: could not render string %s to buffer. %v",
						stringRepr,
						err,
					)
				}
			} else {
				return "", fmt.Errorf(
					"asciiprinter: could not map symbol %d to string. %w",
					symbol,
					UnmappedSymbolErr,
				)
			}
		}

		sb.WriteRune(eol)
	}

	return sb.String(), nil
}

func (asciiPrinter *AsciiPrinter) MapSymbol(symbol CellRenderKind, stringRepr string) {
	asciiPrinter.symbolsMap[symbol] = stringRepr
}

func asSymbol(cell *board.Cell) CellRenderKind {
	if cell.Status(board.Flagged) {
		return Flagged
	}

	if cell.Status(board.Closed) {
		return Closed
	}

	if cell.Mined() {
		return OpenMine
	}

	switch cell.AdjacentMines() {
	case 1:
		return OpenSafe1
	case 2:
		return OpenSafe2
	case 3:
		return OpenSafe3
	case 4:
		return OpenSafe4
	case 5:
		return OpenSafe5
	case 6:
		return OpenSafe6
	case 7:
		return OpenSafe7
	case 8:
		return OpenSafe8
	default:
		return OpenSafe
	}
}

func insertAt(slice []CellRenderKind, index int, item CellRenderKind) []CellRenderKind {
	if index < len(slice) {
		slice[index] = item
		return slice
	}

	if index >= len(slice) {
		newSlice := make([]CellRenderKind, index+1)

		copy(newSlice, slice)

		newSlice[index] = item

		return newSlice
	}

	return slice
}
