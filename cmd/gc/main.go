package main

import (
	"os"

	"github.com/federico-paolillo/mines/internal/gc"
	"github.com/federico-paolillo/mines/internal/runner"
)

func main() {
	statusCode := runner.Run(gc.Program)

	if statusCode == runner.NotOk {
		os.Exit(1)
	}
}
