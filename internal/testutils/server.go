package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/federico-paolillo/mines/internal/server"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

type MinesTestServer struct {
	*http.Server
	Mines *mines.Mines
}

type CustomizeMines = func(*mines.Mines)

func NewCustomServer(
	customizer CustomizeMines,
) *MinesTestServer {
	cfg := &config.Root{
		Seed: 1234,
		Server: config.Server{
			Host: "",
			Port: "65000",
		},
	}

	mines, err := mines.NewMines(
		slog.Default(),
		cfg,
	)
	if err != nil {
		panic(err)
	}

	customizer(mines)

	s := server.NewServer(
		mines,
		cfg.Server,
	)

	return &MinesTestServer{
		s,
		mines,
	}
}

func NewServer() *MinesTestServer {
	return NewCustomServer(
		func(m *mines.Mines) {
		},
	)
}

func NewRequest(
	method string,
	target string,
	body any,
) *http.Request {
	var bodyReader io.Reader = nil

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		bodyReader = strings.NewReader(string(jsonBody))
	}

	return httptest.NewRequest(
		method,
		target,
		bodyReader,
	)
}

func Unmarshal[Body any](
	body *bytes.Buffer,
) (*Body, error) {
	var responseDto Body

	bodyBytes := body.Bytes()

	err := json.Unmarshal(bodyBytes, &responseDto)
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return nil, fmt.Errorf(
			"could not unmarshal response. %v",
			err,
		)
	}

	return &responseDto, nil
}
