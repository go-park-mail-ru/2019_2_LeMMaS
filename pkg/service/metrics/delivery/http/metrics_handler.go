package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/metrics"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const MetricsPath = "/metrics"

type MetricsHandler struct {
	e      *echo.Echo
	logger logger.Logger
}

func NewMetricsHandler(e *echo.Echo, logger logger.Logger) *MetricsHandler {
	metrics.InitHandler()
	handler := MetricsHandler{
		e:      e,
		logger: logger,
	}
	e.GET(MetricsPath, echo.WrapHandler(promhttp.Handler()))
	return &handler
}
