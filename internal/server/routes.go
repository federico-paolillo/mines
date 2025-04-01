package server

import (
	"github.com/federico-paolillo/mines/internal/server/handlers"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/gin-gonic/gin"
)

//	@title			Minewsweeper As a Service API
//	@version		1.0
//	@description	All the endpoints necessary to play Minesweeper matches
//	@accept			json
//	@produce		json
// @licence.name BSD 3-Clause License
// @licence.url https://github.com/federico-paolillo/mines/blob/main/LICENSE
// @contact.name federico-paolillo
// @contact.url https://github.com/federico-paolillo/mines/
func setupHandlers(
	mines *mines.Mines,
	e *gin.Engine,
) {
	attachGameRoutes(mines, e)
}

func attachGameRoutes(
	mines *mines.Mines,
	e *gin.Engine,
) {
	e.POST("/match/:matchId/move", handlers.MakeMove(mines))
	e.GET("/match/:matchId", handlers.GetGame(mines))
	e.POST("/match", handlers.NewGame(mines))
}
