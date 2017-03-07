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
			Buckets: []float64{.000001, .000005, .00001, .00005, .0001, .0005, .001, .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
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
