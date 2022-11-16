package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	CommandStartTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "start_command_total",
	})
	CommandHelpTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "help_command_total",
	})
	CommandAddTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "add_command_total",
	})
	CommandReportTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "report_command_total",
	})
	CommandCurrencyTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "currency_command_total",
	})
	CommandLimitTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "limit_command_total",
	})
	CommandDefaultTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "default_command_total",
	})

	SummaryResponseTime = promauto.NewSummary(prometheus.SummaryOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "summary_response_time_seconds",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.9:  0.01,
			0.99: 0.001,
		},
	})
	HistogramResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "ozon",
			Subsystem: "http",
			Name:      "histogram_response_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
			// Buckets: prometheus.ExponentialBucketsRange(0.0001, 2, 16),
		},
		[]string{"code"},
	)
)
