package main

import (
	"context"
	"os"

	"github.com/federico-paolillo/mines/internal/gc"
	"github.com/federico-paolillo/mines/internal/runner"
)

func main() {
	statusCode := runner.RunMany(
		context.Background(),
		gc.Program,
	)

	if statusCode == runner.NotOk {
		os.Exit(1)
	}
}
