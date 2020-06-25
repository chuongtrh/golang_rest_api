package auth

import (
	"demo_api/src/util/logger"
	"github.com/labstack/echo/v4"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	logger.Info("Load group /api/auth")

	g := e.Group("/api/auth")
	g.POST("/login", controller.Login)
	g.GET("/refresh", controller.Refresh)
}
