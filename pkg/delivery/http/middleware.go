package http

import (
	prometheus "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/monitoring"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type CommonMiddlewaresHandler struct {
	Handler
	e      *echo.Echo
	logger logger.Logger
}

func NewCommonMiddlewaresHandler(e *echo.Echo, logger logger.Logger) CommonMiddlewaresHandler {
	handler := CommonMiddlewaresHandler{Handler{}, e, logger}
	e.Use(handler.panicMiddleware)
	e.Use(handler.prometheusMiddleware)
	return handler
}

func (h CommonMiddlewaresHandler) panicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				h.logger.Errorf("panic during request to %v: %v", c.Request().URL.Path, err)
				h.errorWithStatus(c, "internal error", http.StatusInternalServerError)
			}
		}()
		return next(c)
	}
}

func (h CommonMiddlewaresHandler) prometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		var status = c.Response().Status
		if err != nil {
			status = err.(*echo.HTTPError).Code
		}
		prometheus.ApiMetrics.ObserveResponseTime(status, c.Request().Method, c.Path(), time.Since(start).Seconds())
		prometheus.ApiMetrics.IncHitOfResponse(status, c.Request().Method, c.Path())

		if c.Path() == ApiV1UserRegisterPath && status == 200 {
			prometheus.ApiMetrics.IncRegisteredUsers()
		}
		return nil
	}
}
