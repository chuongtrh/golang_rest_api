//+build wireinject

package server

import (
	"demo_api/src/config"
	"demo_api/src/modules/auth"
	"net/http"

	"github.com/labstack/echo/middleware"

	"demo_api/src/modules/user"
	"demo_api/src/utils"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func initUserController(db *gorm.DB) user.Controller {
	wire.Build(user.NewUserController, user.NewUserService, user.NewUserRepository)
	return user.Controller{}
}

func initAuthController(db *gorm.DB) auth.Controller {
	wire.Build(auth.NewAuthController, user.NewUserService, user.NewUserRepository)
	return auth.Controller{}
}

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "Pong")
}
func loadController(db *gorm.DB, e *echo.Echo) {
	e.Group("/api/ping", Ping)

	auth.LoadRoute(e.Group("/api/auth"), initAuthController(db))
	user.LoadRoute(e.Group("/api/user"), initUserController(db))
}

func Start() {

	//Load env
	if err := config.Load(); err != nil {
		log.Error().Msgf("Error getting env, %v", err)
	} else {
		log.Info().Msg("We are getting the env values")
	}

	//Init echo http
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_unix} ${id} ${method} ${uri} ${status} ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	//Init DB Connection
	if db, err := utils.CreateConnectionDB(); err != nil {

	} else {

		//Load controller
		loadController(db, e)

		//Start app
		AppPort := config.Cfg.AppPort
		if config.Cfg.Env == "local" {
			log.Fatal().Err(e.Start("localhost:" + AppPort))
		} else {
			log.Fatal().Err(e.Start(":" + AppPort))
		}
	}
}
