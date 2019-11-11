package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/access"
	httpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

const (
	CSRFTokenHeader = "X-CSRF-Token"
)

type AccessHandler struct {
	httpDelivery.Handler
	csrfUsecase access.CsrfUsecase
	logger      logger.Logger
}

func NewAccessHandler(e *echo.Echo, csrfUsecase access.CsrfUsecase, logger logger.Logger) *AccessHandler {
	handler := AccessHandler{csrfUsecase: csrfUsecase, logger: logger}
	e.Use(handler.CsrfMiddleware)
	e.GET(httpDelivery.ApiV1AccessCSRFPath, handler.HandleGetCSRFToken)
	return &handler
}

func (h *AccessHandler) HandleGetCSRFToken(c echo.Context) error {
	token := ""
	sessionID, err := c.Cookie(httpDelivery.SessionIDCookieName)
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

func (h *AccessHandler) CsrfMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isPrivate := strings.HasPrefix(c.Request().URL.Path, httpDelivery.ApiV1Private)
		method := c.Request().Method
		if !isPrivate || method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions {
			return next(c)
		}

		csrfToken := c.Request().Header.Get(CSRFTokenHeader)
		if csrfToken == "" {
			return c.JSON(http.StatusForbidden, "csrf token required")
		}
		sessionID, err := c.Cookie(httpDelivery.SessionIDCookieName)
		if err != nil {
			return c.JSON(http.StatusForbidden, "no session cookie")
		}
		ok, err := h.csrfUsecase.CheckTokenBySession(csrfToken, sessionID.Value)
		if !ok {
			message := fmt.Sprintf("recieved incorrect CSRF token %v, session id %v", csrfToken, sessionID.Value)
			if err != nil {
				message += "\n" + err.Error()
			}
			h.logger.Warnf(message)
			return c.JSON(http.StatusForbidden, "incorrect CSRF token")
		}
		return next(c)
	}
}
