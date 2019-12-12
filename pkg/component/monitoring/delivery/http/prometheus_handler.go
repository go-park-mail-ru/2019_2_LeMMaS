package http

import (
	prometheus "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/monitoring"
	delivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusHandler struct {
	delivery.Handler
	e      *echo.Echo
	logger logger.Logger
}

func NewPrometheusHandler(e *echo.Echo, logger logger.Logger) *PrometheusHandler {
	prometheus.InitHandler()
	handler := PrometheusHandler{
		Handler: delivery.Handler{},
		e:       e,
		logger:  logger,
	}
	e.GET(delivery.ApiV1PrometheusPath, echo.WrapHandler(promhttp.Handler()))
	return &handler
}
