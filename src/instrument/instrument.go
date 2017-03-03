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

type options struct {
	port string
}

type operator func(*options)

func Port(port string) operator {
	return func(o *options) {
		o.port = port
	}
}

type instrument struct {
	opts options
}

func NewInstrument(operators ...operator) *instrument {

	var opts options
	opts.port = ":8090"

	for _, o := range operators {
		o(&opts)
	}

	ins := &instrument{
		opts: opts,
	}

	prometheus.MustRegister(RequestDuration)

	return ins
}

func (ins *instrument) Serve() error {

	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(ins.opts.port, nil)
	if err != nil {
		return err
	}

	return nil
}

func NewTimer(api string) *prometheus.Timer {

	return prometheus.NewTimer(
		prometheus.ObserverFunc(func(v float64) {
			RequestDuration.WithLabelValues(api).Observe(v)
		}))
}
