package http

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/access"
	httpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/delivery/http"
	"github.com/labstack/echo"
	"net/http"
)

const (
	ApiV1GetCSRFTokenPath = httpDelivery.ApiV1PathPrefix + "/access/csrf/token"
)

const (
	CSRFTokenHeader = "X-CSRF-Token"
)

type AccessHandler struct {
	csrfUsecase access.CsrfUsecase
	httpDelivery.Handler
}

func NewAccessHandler(e *echo.Echo, csrfUsecase access.CsrfUsecase) *AccessHandler {
	handler := AccessHandler{csrfUsecase: csrfUsecase}
	e.Use(handler.CsrfMiddleware)
	e.GET(ApiV1GetCSRFTokenPath, handler.HandleGetCSRFToken)
	return &handler
}

func (h *AccessHandler) HandleGetCSRFToken(c echo.Context) error {
	token := ""
	tokenType := ""
	sessionID, err := c.Cookie(httpDelivery.SessionIDCookieName)
	if err != nil {
		token, err = h.csrfUsecase.CreateSimpleToken()
		tokenType = "simple"
	} else {
		token, err = h.csrfUsecase.CreateTokenBySession(sessionID.Value)
		tokenType = "session"
	}
	if err != nil {
		c.Logger().Error(err)
		return h.Error(c, errors.New("error generating token"))
	}
	return h.OkWithBody(c, map[string]string{
		"token": token,
		"type":  tokenType,
	})
}

func (h *AccessHandler) CsrfMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		method := c.Request().Method
		if method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions {
			return next(c)
		}
		csrfToken := c.Request().Header.Get(CSRFTokenHeader)
		if csrfToken == "" {
			return c.JSON(http.StatusForbidden, "csrf token required")
		}
		sessionID, err := c.Cookie(httpDelivery.SessionIDCookieName)
		var ok bool
		if err != nil {
			ok, err = h.csrfUsecase.CheckSimpleToken(csrfToken)
		} else {
			ok, err = h.csrfUsecase.CheckTokenBySession(csrfToken, sessionID.Value)
		}
		if !ok {
			c.Echo().Logger.Errorf("recieved incorrect CSRF token %v, session id %v", csrfToken, sessionID.Value)
			if err != nil {
				c.Echo().Logger.Error(err)
			}
			return c.JSON(http.StatusForbidden, "incorrect CSRF token")
		}
		return next(c)
	}
}
