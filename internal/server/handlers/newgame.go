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

// NewGame starts a new Match
//
//	@id			new-match
//	@summary	Starts a new Match that will last roughly 2h
//	@router		/match [post]
//	@param		request	body		req.NewGameDto		true	"Match configuration"
//	@success	200		{object}	res.MatchstateDto	"Updated Match state"
//	@failure	400		"Match configuration format is not correct"
//	@failure	500		"Something went horribly wrong when making the new Match"
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
