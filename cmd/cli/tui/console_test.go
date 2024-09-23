package tui_test

import (
	"io"
	"strings"
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui"
)

func TestConsolePrintLines(t *testing.T) {
	var stdout strings.Builder
	var stdin io.Reader

	c := tui.NewConsole(stdin, &stdout)

	c.Printline("pippo")
	c.Printline("pluto")
	c.Printline("topolino")

	expectedScreen := "pippo\npluto\ntopolino\n"

	screen := stdout.String()

	if expectedScreen != screen {
		t.Fatalf(
			"wrong screen output. wanted '%q' got '%q'",
			expectedScreen,
			screen,
		)
	}
}

func TestConsoleReadLines(t *testing.T) {
	var stdout strings.Builder

	stdin := strings.NewReader(
		"1\n1234\ny\n",
	)

	c := tui.NewConsole(stdin, &stdout)

	lines := c.Scanline()

	expectedLines := "1"

	if expectedLines != lines {
		t.Fatalf(
			"wrong screen input. wanted '%q' got '%q'",
			expectedLines,
			lines,
		)
	}
}
