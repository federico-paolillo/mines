package handlers

import (
	"net/http"

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
		ctx.Status(http.StatusTeapot)
	}
}
