package server_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/server"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

func TestServerReturnsBogusResponseForGetMatch(t *testing.T) {
	cfg := config.Root{
		Seed: 1234,
		Server: config.Server{
			Host: "",
			Port: "65000",
		},
	}

	mines := mines.NewMines(
		slog.Default(),
		cfg,
	)

	s := server.NewServer(
		mines,
		cfg.Server,
	)

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/match/1234", nil)

	s.Handler.ServeHTTP(w, req)

	if w.Code != http.StatusTeapot {
		t.Fatalf(
			"unexpected status code. got %d wanted %d",
			w.Code,
			http.StatusTeapot,
		)
	}
}
