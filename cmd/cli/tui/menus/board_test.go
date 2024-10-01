package menus_test

import (
	"os"
	"strings"
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
	"github.com/federico-paolillo/mines/internal/generators"
	"github.com/federico-paolillo/mines/internal/printers"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

func TestBoardViewRendersProperly(t *testing.T) {
	var stdout strings.Builder

	c := console.NewConsole(
		os.Stdin,
		&stdout,
	)

	bg := generators.NewRngBoardGenerator(1234)

	b := bg.Generate(dimensions.Size{Width: 2, Height: 2}, 1)

	p := printers.NewAsciiPrinter()

	v := menus.NewBoardView(
		c,
		b,
		p,
	)

	v.Render()

	screen := stdout.String()
	screenExpected := "oo\noo\n\n"

	if screen != screenExpected {
		t.Errorf("dialog did not render expected output. wanted '%q' got '%q'", screenExpected, screen)
	}
}
