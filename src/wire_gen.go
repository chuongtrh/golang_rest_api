// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"demo_api/src/config"
	"demo_api/src/modules/auth"
	"demo_api/src/modules/user"
	"demo_api/src/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog/log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Injectors from server.go:

func initUserController(db *gorm.DB) user.Controller {
	repository := user.NewUserRepository(db)
	service := user.NewUserService(repository)
	controller := user.NewUserController(service)
	return controller
}

func initAuthController(db *gorm.DB) auth.Controller {
	repository := user.NewUserRepository(db)
	service := user.NewUserService(repository)
	controller := auth.NewAuthController(service)
	return controller
}

// server.go:

func loadController(db *gorm.DB, e *echo.Echo) {
	auth.LoadRoute(e.Group("/api/auth"), initAuthController(db))
	user.LoadRoute(e.Group("/api/user"), initUserController(db))
}

func Start() {

	if err := config.Load(); err != nil {
		log.Error().Msgf("Error getting env, %v", err)
	} else {
		log.Info().Msg("We are getting the env values")
	}

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_unix} ${id} ${method} ${uri} ${status} ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	if db, err := utils.CreateConnectionDB(); err != nil {

	} else {

		loadController(db, e)

		AppPort := config.Cfg.AppPort
		if config.Cfg.Env == "local" {
			log.Fatal().Err(e.Start("localhost:" + AppPort))
		} else {
			log.Fatal().Err(e.Start(":" + AppPort))
		}
	}
}
