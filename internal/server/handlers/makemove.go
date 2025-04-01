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

// MakeMove applies a Move to an existing Match
//
//	@id			make-move
//	@summary	Applies a Move to an existing Match
//	@router		/match/{matchId}/move [post]
//	@param		matchId	path		string				true	"Match identifier"
//	@param		request	body		req.MoveDto			true	"Move to apply"
//	@success	200		{object}	res.MatchstateDto	"Updated Match state"
//	@failure	400		"Move format is not correct"
//	@failure	404		"Match does not exist"
//	@failure	409		"Your update was superseded by another concurrent update"
//	@failure	422		"The Match has concluded"
//	@failure	500		"Something went horribly wrong when applying the Move to the Match"
func MakeMove(mines *mines.Mines) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		matchId := ctx.Param(req.MatchIdParameterName)

		var moveDto req.MoveDto

		err := ctx.ShouldBindJSON(&moveDto)
		if err != nil {
			mines.Logger.ErrorContext(
				ctx,
				"make move: failed to bind payload",
				slog.Any("err", err),
			)

			ctx.Status(http.StatusBadRequest)

			return
		}

		matchstate, err := mines.Matchmaker.Apply(
			matchId,
			matchmaking.Move{
				Type: moveDto.Type,
				X:    moveDto.X,
				Y:    moveDto.Y,
			},
		)

		if errors.Is(err, matchmaking.ErrNoSuchMatch) {
			ctx.Status(http.StatusNotFound)

			return
		}

		if errors.Is(err, matchmaking.ErrGameHasEnded) {
			ctx.Status(http.StatusUnprocessableEntity)

			return
		}

		if errors.Is(err, matchmaking.ErrConcurrentUpdate) {
			ctx.Status(http.StatusConflict)

			return
		}

		if err != nil {
			mines.Logger.ErrorContext(
				ctx,
				"make game: failed to apply move to match",
				slog.Any("match_id", matchId),
				slog.Any("move_type", moveDto.Type),
				slog.Int("move_x", moveDto.X),
				slog.Int("move_y", moveDto.Y),
				slog.Any("err", err),
			)

			ctx.Status(http.StatusInternalServerError)

			return
		}

		matchstateDto := res.ToMatchstateDto(matchstate)

		ctx.JSON(http.StatusOK, matchstateDto)
	}
}
