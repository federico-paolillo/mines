package console

import (
	"bufio"
	"fmt"
	"io"
)

type Console struct {
	Stdin   io.Reader
	Stdout  io.Writer
	scanner *bufio.Scanner
}

func NewConsole(
	stdin io.Reader,
	stdout io.Writer,
) *Console {
	scanner := bufio.NewScanner(stdin)

	return &Console{
		stdin,
		stdout,
		scanner,
	}
}

func (c *Console) Printline(line string) {
	fmt.Fprintln(c.Stdout, line)
}

func (c *Console) Scanline() string {
	if c.scanner.Scan() {
		return c.scanner.Text()
	}

	return ""
}
