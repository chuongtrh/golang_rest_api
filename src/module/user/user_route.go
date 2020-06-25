package user

import (
	"demo_api/src/util/logger"
	"github.com/labstack/echo/v4"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	logger.Info("Load group /api/user")
	g := e.Group("/api/user")
	g.Use()
	g.GET("/me", controller.GetMyProfile)
	g.GET("", controller.GetAll)
	g.GET("/:id", controller.GetUser)
	g.POST("/register", controller.Register)
}
