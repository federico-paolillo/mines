package handlers

import (
	"net/http"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/gin-gonic/gin"
)

func GetGame(mines *mines.Mines) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusTeapot)
	}
}

func NewGame(mines *mines.Mines) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newGameDto req.NewGameDto

		err := ctx.ShouldBindJSON(&newGameDto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		matchstate, err := mines.Matchmaker.New(newGameDto.Difficulty)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)

			return
		}

		matchstateDto := res.ToMatchstateDto(matchstate)

		ctx.JSON(http.StatusOK, matchstateDto)
	}
}
