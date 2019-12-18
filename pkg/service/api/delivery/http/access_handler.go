package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/labstack/echo"
	"net/http"
	"regexp"
	"strings"
)

const (
	csrfTokenHeader = "X-CSRF-Token"
)

type AccessHandler struct {
	delivery.Handler
	csrf api.CsrfUsecase
	log  logger.Logger
}

func NewAccessHandler(e *echo.Echo, csrf api.CsrfUsecase, log logger.Logger) *AccessHandler {
	handler := AccessHandler{csrf: csrf, log: log}
	e.Use(handler.csrfMiddleware)
	e.Use(handler.corsMiddleware)
	e.GET(delivery.ApiV1AccessCSRFPath, handler.handleGetCSRFToken)
	return &handler
}

func (h *AccessHandler) csrfMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isPrivate := delivery.IsPrivatePath(c.Request().URL.Path)
		method := c.Request().Method
		if !isPrivate || method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions {
			return next(c)
		}

		csrfToken := c.Request().Header.Get(csrfTokenHeader)
		if csrfToken == "" {
			return h.Error(c, "csrf token required")
		}
		session, err := c.Cookie(delivery.SessionCookieName)
		if err != nil {
			return h.Error(c, "no session cookie")
		}
		ok, err := h.csrf.CheckTokenBySession(csrfToken, session.Value)
		if !ok {
			h.warnIncorrectCSRF(csrfToken, c.Path(), err)
			return h.Error(c, "incorrect CSRF token")
		}
		return next(c)
	}
}

func (h *AccessHandler) warnIncorrectCSRF(token, path string, err error) {
	message := fmt.Sprintf("recieved incorrect CSRF token %v, path %v", token, path)
	if err != nil {
		message += "\n" + err.Error()
	}
	h.log.Warnf(message)
}

func (h *AccessHandler) corsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origin := c.Request().Header.Get(echo.HeaderOrigin)
		var allowOrigin string
		if h.isOriginAllowed(origin) {
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

func (h *AccessHandler) isOriginAllowed(origin string) bool {
	isNowSh, _ := regexp.MatchString(`^https:\/\/20192lemmas.*\.now\.sh$`, origin)
	isLocalhost, _ := regexp.MatchString(`^http:\/\/localhost:\d*$`, origin)
	return isNowSh || isLocalhost
}

func (h *AccessHandler) handleGetCSRFToken(c echo.Context) error {
	token := ""
	session, err := c.Cookie(delivery.SessionCookieName)
	if err != nil {
		return h.Error(c, "no session cookie")
	}
	token, err = h.csrf.CreateTokenBySession(session.Value)
	if err != nil {
		h.log.Error(err)
		return h.Error(c, "error generating token")
	}
	return h.OkWithBody(c, map[string]string{
		"token": token,
	})
}
