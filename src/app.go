package app

import (
	"context"
	"demo_api/src/config"
	auth_middleware "demo_api/src/middleware"
	"demo_api/src/module"
	"demo_api/src/util"

	"github.com/casbin/casbin/v2"
	//casbin_mw "github.com/labstack/echo-contrib/casbin"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	enforcer, err := casbin.NewEnforcer("src/config/auth_model.conf", "src/config/auth_policy.csv")
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	enforcer.EnableLog(true)
	enforcer.EnableEnforce(true)

	e.Use(auth_middleware.Authorizer(enforcer))

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

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
		log.Panic().Msgf("Error getting env, %v", err)
	} else {
		log.Info().Msg("Load env Ok!")

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
