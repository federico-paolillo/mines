package middlewares

import (
	"log/slog"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(mines *mines.Mines) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mines.Logger.InfoContext(
			ctx,
			"gin: http request started",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.String()),
		)

		ctx.Next()

		mines.Logger.InfoContext(
			ctx,
			"gin: http request completed",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.String()),
		)
	}
}
