package http

import (
	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func InitMiddlewares(e *echo.Echo) {
	e.Use(corsMiddleware)
	e.Use(panicMiddleware)
	e.Use(middleware.LoggerWithConfig(setLoggerConfig()))
}

func setLoggerConfig() middleware.LoggerConfig {
	f, err := os.OpenFile("agario.log",
		os.O_RDWR | os.O_CREATE | os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()

	return middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","file":"${long_file}"` + "\n",
		CustomTimeFormat: "2000-01-01 15:01:02",
		Output: f,
	}
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

		//println(c.Logger)

		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusNoContent)
		}
		return next(c)
	}
}

func isOriginAllowed(origin string) bool {
	isNowSh, _ := regexp.MatchString(`^https:\/\/20192lemmas-.*\.now\.sh$`, origin)
	isLocalhost, _ := regexp.MatchString(`^http:\/\/localhost:\d*$`, origin)
	return isNowSh || isLocalhost
}

func panicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic during request to %s: %s", c.Request().URL.Path, err)
				c.JSON(http.StatusInternalServerError, "internal error")
			}
		}()
		return next(c)
	}
}
