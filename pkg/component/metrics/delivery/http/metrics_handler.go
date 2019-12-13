package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/metrics"
	delivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsHandler struct {
	delivery.Handler
	e *echo.Echo
}

func NewMetricsHandler(e *echo.Echo) *MetricsHandler {
	metrics.InitHandler()
	h := MetricsHandler{e: e}
	e.GET(delivery.MetricsPath, echo.WrapHandler(promhttp.Handler()))
	return &h
}
