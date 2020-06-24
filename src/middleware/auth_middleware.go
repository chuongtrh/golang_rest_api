package middleware

import (
	"demo_api/src/config"
	"demo_api/src/util"
	"github.com/casbin/casbin/v2"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func getJWTFromHeader(c echo.Context, header string, authScheme string) (string, bool) {
	auth := c.Request().Header.Get(header)
	l := len(authScheme)
	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], true
	}
	return "", false
}

func checkAuthenticate(c echo.Context) (string, error) {
	token, isExist := getJWTFromHeader(c, "Authorization", "Bearer")
	if !isExist {
		return "*", nil
	}
	claims := &util.Claims{}
	err := util.DecodeToken(token, claims, config.Cfg.JwtKey)
	if err != nil {
		return "*", err
	}
	c.Set("user", claims)
	log.Info().Msgf("claims:%+v", claims)
	return claims.Role, nil
}

// CheckPermission func
func checkPermission(ce *casbin.Enforcer, c echo.Context) (bool, error) {
	role, err := checkAuthenticate(c)
	if err != nil {
		log.Error().Msgf(err.Error())
		return false, err
	}

	method := c.Request().Method
	path := c.Request().URL.Path
	return ce.Enforce(role, path, method)
}

// Authorizer func
func Authorizer(ce *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if pass, err := checkPermission(ce, c); err == nil && pass {
				return next(c)
			} else if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			return echo.ErrForbidden
		}
	}
}
