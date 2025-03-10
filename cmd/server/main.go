package main

import (
	"context"
	"flag"
	"os"

	"github.com/federico-paolillo/mines/internal/gc"
	"github.com/federico-paolillo/mines/internal/runner"
	"github.com/federico-paolillo/mines/internal/server"
)

func main() {
	withGc := flag.Bool("gc", false, "Embed garbage collection into server process")

	flag.Parse()

	var statusCode runner.StatusCode

	if *withGc {
		statusCode = runner.RunMany(
			context.Background(),
			server.Program,
			gc.Program,
		)
	} else {
		statusCode = runner.RunMany(
			context.Background(),
			server.Program,
		)
	}

	if statusCode == runner.NotOk {
		os.Exit(1)
	}
}
