package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/labstack/echo"
)

type Middleware struct {
	delivery.Handler
	e      *echo.Echo
	logger logger.Logger
}

func NewMiddleware(e *echo.Echo, logger logger.Logger) Middleware {
	handler := Middleware{delivery.Handler{}, e, logger}
	e.Use(handler.panicMiddleware)
	return handler
}

func (h Middleware) panicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				h.logger.Errorf("panic during request to %v: %v", c.Request().URL.Path, err)
				h.InternalError(c, "internal error")
			}
		}()
		return next(c)
	}
}
