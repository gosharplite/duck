package instrument

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Histogram for request duration, partitioned by API.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"api"},
	)
)

func init() {

	prometheus.MustRegister(RequestDuration)

	http.Handle("/metrics", promhttp.Handler())
}

func NewTimer(api string) *prometheus.Timer {

	return prometheus.NewTimer(
		prometheus.ObserverFunc(func(v float64) {
			RequestDuration.WithLabelValues(api).Observe(v)
		}))
}
