package main

import (
	"context"
	"os"

	"github.com/federico-paolillo/mines/internal/runner"
	"github.com/federico-paolillo/mines/internal/server"
)

func main() {
	statusCode := runner.RunMany(
		context.Background(),
		server.Program,
	)

	if statusCode == runner.NotOk {
		os.Exit(1)
	}
}
