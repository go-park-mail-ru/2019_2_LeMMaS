package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/logger"
	userHttpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/user/delivery/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"regexp"
	"strings"
)

const CSRFTokenHeader = "X-CSRF-Token"
const IncorrectCSRFTokenMessage = "incorrect CSRF token"

func InitMiddlewares(e *echo.Echo) {
	e.Use(corsMiddleware)
	e.Use(panicMiddleware)
	e.Use(csrfMiddleware)
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

func csrfMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		method := c.Request().Method
		if method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions {
			return next(c)
		}
		csrfToken := c.Request().Header.Get(CSRFTokenHeader)
		if csrfToken == "" {
			return c.JSON(http.StatusForbidden, IncorrectCSRFTokenMessage)
		}
		sessionID, err := c.Cookie(userHttpDelivery.SessionIDCookieName)
		if err != nil {
			return c.JSON(http.StatusForbidden, IncorrectCSRFTokenMessage)
		}
		ok, _ := checkCSRFToken(sessionID.Value, csrfToken)
		if !ok {
			// todo: log error
			return c.JSON(http.StatusForbidden, IncorrectCSRFTokenMessage)
		}
		return next(c)
	}
}
