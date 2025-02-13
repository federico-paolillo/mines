package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/gin-gonic/gin"
)

func GetGame(mines *mines.Mines) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		matchId := ctx.Param(req.MatchIdParameterName)

		matchstate, err := mines.Matchmaker.Get(matchId)

		if errors.Is(err, matchmaking.ErrNoSuchMatch) {
			ctx.Status(http.StatusNotFound)

			return
		}

		if err != nil {
			mines.Logger.ErrorContext(
				ctx,
				"get game: failed to retrieve match",
				slog.Any("match_id", matchId),
				slog.Any("err", err),
			)

			ctx.Status(http.StatusInternalServerError)

			return
		}

		matchstateDto := res.ToMatchstateDto(matchstate)

		ctx.JSON(http.StatusOK, matchstateDto)
	}
}
