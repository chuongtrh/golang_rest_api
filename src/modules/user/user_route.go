package user

import (
	"demo_api/src/middlewares"

	"github.com/labstack/echo"
)

func LoadRoute(g *echo.Group, controller Controller) {
	g.GET("/me", controller.GetMyProfile, middlewares.IsAuthenticate)
	g.GET("", controller.GetAll, middlewares.CheckPermission([]string{RoleAdmin}))
	g.GET("/:id", controller.GetUser, middlewares.CheckPermission([]string{RoleAdmin}))
}
