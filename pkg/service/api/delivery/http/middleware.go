package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/labstack/echo"
	"time"
)

type Middleware struct {
	delivery.Handler
	metrics api.Metrics
	e       *echo.Echo
	log     logger.Logger
}

func NewMiddleware(e *echo.Echo, metrics api.Metrics, log logger.Logger) Middleware {
	handler := Middleware{
		Handler: delivery.Handler{},
		metrics: metrics,
		e:       e,
		log:     log,
	}
	e.Use(handler.panicMiddleware)
	e.Use(handler.metricsMiddleware)
	return handler
}

func (h Middleware) panicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				h.log.Errorf("panic during request to %v: %v", c.Request().URL.Path, err)
				h.InternalError(c, "internal error")
			}
		}()
		return next(c)
	}
}

func (h Middleware) metricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		var status = c.Response().Status
		if err != nil {
			status = err.(*echo.HTTPError).Code
		}
		h.metrics.ObserveResponseTime(status, c.Request().Method, c.Path(), time.Since(start).Seconds())
		h.metrics.IncHits(status, c.Request().Method, c.Path())
		return nil
	}
}
