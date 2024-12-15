package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

var (
    requestCount = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "user_service_requests_total",
            Help: "Total number of requests received by the user service",
        },
        []string{"method"},
    )

    dbConnectionCount = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "user_service_db_connections",
            Help: "Current number of database connections",
        },
    )
)

func init() {
    prometheus.MustRegister(requestCount)
    prometheus.MustRegister(dbConnectionCount)
}

func RecordRequest(method string) {
    requestCount.WithLabelValues(method).Inc()
}

func SetDBConnectionCount(count float64) {
    dbConnectionCount.Set(count)
}

func StartMetricsServer(addr string) {
    http.Handle("/metrics", promhttp.Handler())
    go func() {
        if err := http.ListenAndServe(addr, nil); err != nil {
            // Handle error (e.g., log it)
        }
    }()
}