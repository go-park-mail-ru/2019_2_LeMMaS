package http

import (
	"errors"
	userHttpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/user/delivery/http"
	"github.com/labstack/echo"
)

const ApiV1GetCSRFTokenPath = ApiV1PathPrefix + "/csrf/token"

type CSRFHandler struct {
	Handler
}

func NewCSRFHandler(e *echo.Echo) *CSRFHandler {
	handler := CSRFHandler{}
	e.GET(ApiV1GetCSRFTokenPath, handler.HandleGetToken)
	return &handler
}

func (h *CSRFHandler) HandleGetToken(c echo.Context) error {
	sessionID, err := c.Cookie(userHttpDelivery.SessionIDCookieName)
	if err != nil {
		return h.Error(c, errors.New("no session cookie"))
	}
	token, err := createCSRFToken(sessionID.Value)
	if err != nil {
		return h.Error(c, errors.New("error generating token"))
	}
	return h.OkWithBody(c, map[string]string{
		"token": token,
	})
}
