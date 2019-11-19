package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/labstack/echo"
	"net/http"
)

type CommonMiddlewaresHandler struct {
	Handler
	e      *echo.Echo
	logger logger.Logger
}

func NewCommonMiddlewaresHandler(e *echo.Echo, logger logger.Logger) CommonMiddlewaresHandler {
	handler := CommonMiddlewaresHandler{Handler{}, e, logger}
	e.Use(handler.panicMiddleware)
	return handler
}

func (h CommonMiddlewaresHandler) panicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				h.logger.Errorf("panic during request to %v: %v", c.Request().URL.Path, err)
				h.ErrorWithStatus(c, "internal error", http.StatusInternalServerError)
			}
		}()
		return next(c)
	}
}
