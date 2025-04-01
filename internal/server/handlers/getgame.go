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

// GetGame retrieves an existing Match
//
//	@id			get-match
//	@summary	Gets a Match with the identifier specified
//	@router		/match/{matchId} [get]
//	@param		matchId	path		string				true	"Match identifier"
//	@success	200		{object}	res.MatchstateDto	"Current Match state"
//	@failure	404		"Match does not exist"
//	@failure	500		"Something went horribly wrong when retrieving the Match"
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
