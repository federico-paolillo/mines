package main

import (
	"os"

	"github.com/federico-paolillo/mines/internal/runner"
	"github.com/federico-paolillo/mines/internal/server"
)

func main() {
	statusCode := runner.Run(server.Program)

	if statusCode == runner.NotOk {
		os.Exit(1)
	}
}
