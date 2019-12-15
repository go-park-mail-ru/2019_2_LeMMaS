package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	prometheus "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/monitoring"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const MetricsPath = "/metrics"

type PrometheusHandler struct {
	e      *echo.Echo
	logger logger.Logger
}

func NewPrometheusHandler(e *echo.Echo, logger logger.Logger) *PrometheusHandler {
	prometheus.InitHandler()
	handler := PrometheusHandler{
		e:      e,
		logger: logger,
	}
	e.GET(MetricsPath, echo.WrapHandler(promhttp.Handler()))
	return &handler
}
