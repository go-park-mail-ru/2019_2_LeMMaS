package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type Metrics interface {
	IncHits(status int, method, path string)
	ObserveResponseTime(status int, method, path string, observeTime float64)
}

type metrics struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
}

func NewMetrics() (Metrics, error) {
	m := metrics{}
	m.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hits_total",
	})
	m.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hits",
		},
		[]string{"status", "method", "path"},
	)
	m.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "times",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(m.HitsTotal); err != nil {
		return nil, err
	}
	if err := prometheus.Register(m.Hits); err != nil {
		return nil, err
	}
	if err := prometheus.Register(m.Times); err != nil {
		return nil, err
	}
	if err := prometheus.Register(prometheus.NewBuildInfoCollector()); err != nil {
		return nil, err
	}

	return &m, nil
}

func (m *metrics) IncHits(status int, method, path string) {
	m.HitsTotal.Inc()
	m.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

func (m *metrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	m.Times.WithLabelValues(strconv.Itoa(status), method, path).Observe(observeTime)
}
