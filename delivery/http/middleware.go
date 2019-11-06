package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"regexp"
	"strings"
)

func InitMiddlewares(e *echo.Echo) {
	e.Use(corsMiddleware)
	e.Use(panicMiddleware)
	e.Use(middleware.LoggerWithConfig(logger.SetLoggerConfig()))
}

func corsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origin := c.Request().Header.Get(echo.HeaderOrigin)
		var allowOrigin string
		if isOriginAllowed(origin) {
			allowOrigin = origin
		} else {
			allowOrigin = ""
		}
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)

		allowedMethods := []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete}
		c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, strings.Join(allowedMethods, ","))
		c.Response().Header().Set(echo.HeaderAccessControlAllowCredentials, "true")

		allowedHeaders := []string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}
		c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, strings.Join(allowedHeaders, ","))

		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusNoContent)
		}
		return next(c)
	}
}

func isOriginAllowed(origin string) bool {
	isNowSh, _ := regexp.MatchString(`^https:\/\/20192lemmas.*\.now\.sh$`, origin)
	isLocalhost, _ := regexp.MatchString(`^http:\/\/localhost:\d*$`, origin)
	return isNowSh || isLocalhost
}

func panicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.Echo().Logger.Printf("panic during request to %s: %s", c.Request().URL.Path, err)
				c.JSON(http.StatusInternalServerError, "internal error")
			}
		}()
		return next(c)
	}
}
