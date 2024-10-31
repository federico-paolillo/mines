package server

import (
	"net/http"
	"time"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/gin-gonic/gin"
)

func NewServer(
	mines *mines.Mines,
	cfg config.Server,
) *http.Server {
	e := gin.New()

	setupMiddlewares(mines, e)
	setupHandlers(mines, e)

	s := &http.Server{
		Addr:         cfg.Endpoint(),
		Handler:      e,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}

	return s
}

func setupMiddlewares(
	mines *mines.Mines,
	e *gin.Engine,
) {
	// TODO: Add authz middleware
}

func setupHandlers(
	mines *mines.Mines,
	e *gin.Engine,
) {
	attachGameRoutes(mines, e)
}
