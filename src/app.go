package app

import (
	"context"
	"demo_api/src/config"
	"demo_api/src/module"
	"demo_api/src/util"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func initEcho() *echo.Echo {
	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_unix} ${id} ${method} ${uri} ${status} ${latency_human}\n",
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	e.Use(middleware.CSRF())
	e.Use(middleware.CORS())

	return e
}

func register(lc fx.Lifecycle, e *echo.Echo, db *gorm.DB) {
	log.Info().Msg("Register.")

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Info().Msg("Starting server.")

				e.GET("/api/ping", func(c echo.Context) error {
					return c.JSON(http.StatusOK, "Pong")
				})

				AppPort := config.Cfg.AppPort
				if config.Cfg.Env == "local" {
					go e.Start("localhost:" + AppPort)
				} else {
					go e.Start(":" + AppPort)
				}
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info().Msg("Shutting down server.")
				_ = db.Close()
				return e.Close()
			},
		},
	)
}

// Run server
func Run() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	//Load env
	if err := config.Load(); err != nil {
		log.Error().Msgf("Error getting env, %v", err)
	} else {
		log.Info().Msg("We are getting the env values")

		ServerDependencies := fx.Provide(
			util.CreateConnectionDB,
			initEcho,
		)

		fx.New(fx.Options(
			ServerDependencies,
			module.Module,
			fx.Invoke(register),
		)).Run()
	}
}
