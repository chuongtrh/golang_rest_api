package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	log.Info().Msg("Load group /api/auth")

	g := e.Group("/api/auth")
	g.POST("/login", controller.Login)
	g.GET("/refresh", controller.Refresh)
}
