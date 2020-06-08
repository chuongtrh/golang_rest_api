package user

import (
	"demo_api/src/middleware"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	log.Info().Msg("Load group /api/user")
	g := e.Group("/api/user")
	g.GET("/me", controller.GetMyProfile, middleware.IsAuthenticate)
	g.GET("", controller.GetAll, middleware.CheckPermission([]string{RoleAdmin}))
	g.GET("/:id", controller.GetUser, middleware.CheckPermission([]string{RoleAdmin}))
	g.POST("/register", controller.Register)
}
