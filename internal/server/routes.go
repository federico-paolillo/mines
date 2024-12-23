package server

import (
	"github.com/federico-paolillo/mines/internal/server/handlers"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/gin-gonic/gin"
)

func attachGameRoutes(
	mines *mines.Mines,
	e *gin.Engine,
) {
	e.POST("/match/:matchId/move", handlers.MakeMove(mines))
	e.GET("/match/:matchId", handlers.GetGame(mines))
	e.POST("/match", handlers.NewGame(mines))
}
