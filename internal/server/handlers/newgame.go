package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/gin-gonic/gin"
)

func NewGame(mines *mines.Mines) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newGameDto req.NewGameDto

		err := ctx.ShouldBindJSON(&newGameDto)
		if err != nil {
			mines.Logger.ErrorContext(
				ctx,
				"new game: failed to bind payload",
				slog.Any("err", err),
			)

			ctx.Status(http.StatusBadRequest)

			return
		}

		matchstate, err := mines.Matchmaker.New(
			time.Now().Unix(),
			newGameDto.Difficulty,
		)
		if err != nil {
			mines.Logger.ErrorContext(
				ctx,
				"new game: failed to create new match",
				slog.Any("difficulty", newGameDto.Difficulty),
				slog.Any("err", err),
			)

			ctx.Status(http.StatusInternalServerError)

			return
		}

		matchstateDto := res.ToMatchstateDto(matchstate)

		ctx.JSON(http.StatusOK, matchstateDto)
	}
}
