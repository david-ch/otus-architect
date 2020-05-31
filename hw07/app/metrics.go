package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	latencyHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "app_request_latency_seconds",
			Help: "Application request latency",
		},
		[]string{
			"method",
			"endpoint",
		},
	)

	requestsCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_request_count",
			Help: "Application request count",
		},
		[]string{
			"method",
			"endpoint",
			"http_status",
		},
	)
)

func registerMetrics(r *mux.Router) {
	r.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(latencyHistogram)
	prometheus.MustRegister(requestsCount)
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}

func withMetrics(endpoint string, handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapperWriter := &statusInterceptingRespWriter{w, 200}

		handler(wrapperWriter, r)
		elapsed := time.Since(start)

		latencyHistogram.With(
			prometheus.Labels{
				"method":   r.Method,
				"endpoint": endpoint,
			},
		).Observe(elapsed.Seconds())

		requestsCount.With(
			prometheus.Labels{
				"method":      r.Method,
				"endpoint":    endpoint,
				"http_status": fmt.Sprintf("%d", wrapperWriter.status),
			},
		).Inc()
	}
}

type statusInterceptingRespWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusInterceptingRespWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
