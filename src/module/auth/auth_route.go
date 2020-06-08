package auth

import (
	"demo_api/src/middleware"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	log.Info().Msg("Load group /api/auth")

	g := e.Group("/api/auth")
	g.POST("/login", controller.Login)
	g.GET("/refresh", controller.Refresh, middleware.IsAuthenticate)
}
