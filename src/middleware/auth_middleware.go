package middleware

import (
	"demo_api/src/config"
	"demo_api/src/util"
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

func jwtFromHeader(c echo.Context, header string, authScheme string) (string, error) {
	auth := c.Request().Header.Get(header)
	l := len(authScheme)
	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], nil
	}
	return "", errors.New("missing authorization")
}

func checkAuthenticate(c echo.Context) error {
	token, err := jwtFromHeader(c, "Authorization", "Bearer")
	if err != nil {
		return err
	}
	claims := &util.Claims{}
	err = util.DecodeToken(token, claims, config.Cfg.JwtKey)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	c.Set("user", claims)

	return nil
}

// IsAuthenticate func
func IsAuthenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := checkAuthenticate(c); err != nil {
			return err
		}
		return next(c)
	}
}

// CheckPermission func
func CheckPermission(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := checkAuthenticate(c); err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}
			claims := c.Get("user").(*util.Claims)
			log.Info().Msgf("claims:%+v", claims)
			if funk.Contains(roles, claims.Role) {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
		}
	}
}
