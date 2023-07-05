package vPrometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type AppPrometheus struct {
	Counter *prometheus.GaugeVec
}

func CreatePrometheus() *AppPrometheus {
	return &AppPrometheus{}
}

func (r *AppPrometheus) InitPrometheus() {

	r.Counter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "system",
			Subsystem: "jobs",
			Name:      "timeouts_total",
			Help:      "Total timeouts of jobs processed by the nodes",
		},
		[]string{"host", "service", "counter"},
	)
	prometheus.MustRegister(r.Counter)
}

func (r *AppPrometheus) WriteMetric(host string, service string, value float64) {
	r.Counter.WithLabelValues(host, service, "counter").Set(float64(value))
}
