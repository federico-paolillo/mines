package menus

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/internal/printers"
	"github.com/federico-paolillo/mines/pkg/board"
)

type BoardView struct {
	board   *board.Board
	printer *printers.AsciiPrinter
	console *console.Console
}

func NewBoardView(
	console *console.Console,
	board *board.Board,
	printer *printers.AsciiPrinter,
) *BoardView {
	return &BoardView{
		board,
		printer,
		console,
	}
}

func (b *BoardView) Render() {
	output, err := b.printer.Render(b.board)

	if err == nil {
		b.console.Printline(output)
	} else {
		b.console.Printline(err.Error())
	}
}
