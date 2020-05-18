package auth

import (
	"demo_api/src/middlewares"

	"github.com/labstack/echo"
)

func LoadRoute(g *echo.Group, controller Controller) {
	g.POST("/login", controller.Login)
	g.GET("/refresh", controller.Refresh, middlewares.IsAuthenticate)
}
