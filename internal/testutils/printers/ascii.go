package printers

import (
	"strconv"
	"strings"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

const eol = '\n'

func Render(board *board.Board) string {
	var sb strings.Builder

	size := board.Size()

	for y := range size.Height {
		for x := range size.Width {
			cell := board.Retrieve(
				dimensions.Location{
					X: x + 1,
					Y: y + 1,
				},
			) // Boards start at 1,1

			symbol := asSymbol(cell)

			_, _ = sb.WriteString(symbol)
		}

		sb.WriteRune(eol)
	}

	return sb.String()
}

func asSymbol(cell *board.Cell) string {
	if cell.HasStatus(board.ClosedCell) {
		return "o"
	}

	if cell.HasStatus(board.FlaggedCell) {
		return "f"
	}

	if cell.Mined() {
		return "x"
	}

	adjacentMines := cell.AdjacentMines()

	return strconv.Itoa(adjacentMines)
}
