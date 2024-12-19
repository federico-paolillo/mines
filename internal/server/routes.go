package server

import (
	"github.com/federico-paolillo/mines/internal/server/handlers"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/gin-gonic/gin"
)

// POST /match 						 => 200 (GameState)
// GET  /match/<uuid>		   => 200 (GameState), 404 (NotFound)
// POST /match/<uuid>/move => 200 (GameState), 400 (Validation), 422 (Game Lost/Won)

func attachGameRoutes(
	mines *mines.Mines,
	e *gin.Engine,
) {
	e.GET("/match/:matchId", handlers.GetGame(mines))
	e.POST("/match", handlers.NewGame(mines))
}
