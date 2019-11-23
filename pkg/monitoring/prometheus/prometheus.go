package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	normDomain        = 0.0002
	normMean          = 0.00001
)

var (
	usersRegistered = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name:		"count of users",
		})
	rpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "rpc_durations_seconds",
			Help:       "RPC latency distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"service"},
	)

	rpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "rpc_durations_histogram_seconds",
		Help:    "RPC latency distributions.",
		Buckets: prometheus.LinearBuckets(normMean-5*normDomain, .5*normDomain, 20),
	})
)

func init() {
	prometheus.MustRegister(usersRegistered)
	prometheus.MustRegister(rpcDurations)
	prometheus.MustRegister(rpcDurationsHistogram)
	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())

	go func() {
		for {
			usersRegistered.Inc()
			time.Sleep(1000 * time.Millisecond)
		}
	}()
}