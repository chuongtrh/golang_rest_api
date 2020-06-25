package middleware

import (
	"demo_api/src/util/logger"
	"github.com/labstack/echo/v4"
	"time"
)

func ZapLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			logger.Infof("%s\t%s\t%s\t%s\t%d\t%s", "REQUEST", id, req.Method, req.RequestURI, res.Status, time.Since(start).String())

			return nil
		}
	}
}
