package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsHandler struct {
}

func NewMetricsHandler(e *echo.Echo) *MetricsHandler {
	handler := MetricsHandler{}
	e.GET(delivery.MetricsPath, echo.WrapHandler(promhttp.Handler()))
	return &handler
}
