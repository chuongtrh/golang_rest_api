package user

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	log.Info().Msg("Load group /api/user")
	g := e.Group("/api/user")
	g.Use()
	g.GET("/me", controller.GetMyProfile)
	g.GET("", controller.GetAll)
	g.GET("/:id", controller.GetUser)
	g.POST("/register", controller.Register)
}
