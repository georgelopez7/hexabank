package metrics

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewtMetricsEndpoint() {
	metricsPort := os.Getenv("METRICS_PORT")
	if metricsPort == "" {
		log.Fatal("METRICS_PORT environment variable is not set")
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":"+metricsPort, nil)
	}()
}
