package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/access"
	delivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
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
	csrfUsecase access.CsrfUsecase
	logger      logger.Logger
}

func NewAccessHandler(e *echo.Echo, csrfUsecase access.CsrfUsecase, logger logger.Logger) *AccessHandler {
	handler := AccessHandler{csrfUsecase: csrfUsecase, logger: logger}
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
		sessionID, err := c.Cookie(delivery.SessionIDCookieName)
		if err != nil {
			return h.Error(c, "no session cookie")
		}
		ok, err := h.csrfUsecase.CheckTokenBySession(csrfToken, sessionID.Value)
		if !ok {
			message := fmt.Sprintf("recieved incorrect CSRF token %v, session id %v", csrfToken, sessionID.Value)
			if err != nil {
				message += "\n" + err.Error()
			}
			h.logger.Warnf(message)
			return h.Error(c, "incorrect CSRF token")
		}
		return next(c)
	}
}

func (h AccessHandler) corsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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

func (h AccessHandler) isOriginAllowed(origin string) bool {
	isNowSh, _ := regexp.MatchString(`^https:\/\/20192lemmas.*\.now\.sh$`, origin)
	isLocalhost, _ := regexp.MatchString(`^http:\/\/localhost:\d*$`, origin)
	return isNowSh || isLocalhost
}

func (h *AccessHandler) handleGetCSRFToken(c echo.Context) error {
	token := ""
	sessionID, err := c.Cookie(delivery.SessionIDCookieName)
	if err != nil {
		return h.Error(c, "no session cookie")
	}
	token, err = h.csrfUsecase.CreateTokenBySession(sessionID.Value)
	if err != nil {
		h.logger.Error(err)
		return h.Error(c, "error generating token")
	}
	return h.OkWithBody(c, map[string]string{
		"token": token,
	})
}
