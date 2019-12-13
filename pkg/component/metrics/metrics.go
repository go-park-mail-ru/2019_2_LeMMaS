package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type Metrics struct {
	RegisteredUsers prometheus.Counter
	HitsTotal       prometheus.Counter
	Hits            *prometheus.CounterVec
	Times           *prometheus.HistogramVec
}

var API Metrics

func initMetrics() {
	API.RegisteredUsers = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "count_of_users"},
	)
	API.HitsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "hits_total"},
	)
	API.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "hits"},
		[]string{"status", "method", "path"},
	)
	API.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{Name: "times"},
		[]string{"status", "method", "path"},
	)
}

func (m *Metrics) IncHitsTotal() {
	m.HitsTotal.Inc()
}

func (m *Metrics) IncHitOfResponse(status int, method, path string) {
	m.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

func (m *Metrics) IncRegisteredUsers() {
	m.RegisteredUsers.Inc()
}

func (m *Metrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	m.Times.WithLabelValues(strconv.Itoa(status), method, path).Observe(observeTime)
}

func InitHandler() {
	initMetrics()
	prometheus.MustRegister(API.HitsTotal)
	prometheus.MustRegister(API.Hits)
	prometheus.MustRegister(API.Times)
	prometheus.MustRegister(API.RegisteredUsers)
	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}
