package server

import (
	"net/http"
	"time"

	"github.com/federico-paolillo/mines/internal/server/middlewares"
	"github.com/federico-paolillo/mines/internal/server/validators"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func NewServer(
	mines *mines.Mines,
	cfg config.Server,
) *http.Server {
	e := gin.New()

	setupValidation()

	setupMiddlewares(mines, e)
	setupHandlers(mines, e)

	e.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusTeapot)
	})

	s := &http.Server{
		Addr:         cfg.Endpoint(),
		Handler:      e,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}

	return s
}

func setupValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation(validators.IsDifficultyEnumValidator, validators.IsDifficultyEnum)
		_ = v.RegisterValidation(validators.IsMoveTypeEnumValidator, validators.IsMoveTypeEnum)
	}
}

func setupMiddlewares(
	mines *mines.Mines,
	e *gin.Engine,
) {
	e.Use(middlewares.LoggingMiddleware(mines))
}

