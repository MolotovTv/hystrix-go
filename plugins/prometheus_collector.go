package plugins

import (
	"time"

	"github.com/afex/hystrix-go/hystrix/metric_collector"
	"github.com/prometheus/client_golang/prometheus"
)

var promAttempts = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_attempts_total",
		Help: "Hytrix attemps",
	},
	[]string{"circuit_name"},
)

var promErrors = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_error_total",
		Help: "Hytrix errors",
	},
	[]string{"circuit_name"},
)

var promFailures = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_failure_total",
		Help: "Hytrix failures",
	},
	[]string{"circuit_name"},
)

var promRejects = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_rejects_total",
		Help: "Hytrix rejects",
	},
	[]string{"circuit_name"},
)

var promShortCircuits = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_short_circuits_total",
		Help: "Hytrix short circuits",
	},
	[]string{"circuit_name"},
)

var promTimeouts = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_timeouts_total",
		Help: "Hytrix timeouts",
	},
	[]string{"circuit_name"},
)

var promFallbackSuccesses = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_fallback_success_total",
		Help: "Hytrix fallback successes",
	},
	[]string{"circuit_name"},
)

var promFallbackFailures = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_fallback_failures_total",
		Help: "Hytrix fallback failures",
	},
	[]string{"circuit_name"},
)

var promSuccesses = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_success_total",
		Help: "Hytrix success",
	},
	[]string{"circuit_name"},
)

var promTotalDuration = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "hystrix_total_duration_seconds",
		Help: "Hystric total duration",
	},
	[]string{"circuit_name"},
)

var promRunDuration = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hystrix_run_duration_seconds",
		Help: "Hystric run duration",
	},
	[]string{"circuit_name"},
)

type PrometheusCollector struct {
	circuitName string
}

func NewPrometheusCollector() func(string) metricCollector.MetricCollector {
	prometheus.MustRegister(
		promAttempts, promErrors, promFailures,
		promRejects, promShortCircuits, promTimeouts,
		promFallbackSuccesses, promFallbackFailures,
		promSuccesses, promTotalDuration, promRunDuration,
	)

	return func(name string) metricCollector.MetricCollector {
		return &PrometheusCollector{
			circuitName: name,
		}
	}
}

// IncrementAttempts increments the number of calls to this circuit.
func (c *PrometheusCollector) IncrementAttempts() {
	promAttempts.WithLabelValues(c.circuitName).Inc()
}

// IncrementErrors increments the number of unsuccessful attempts.
// Attempts minus Errors will equal successes within a time range.
// Errors are any result from an attempt that is not a success.
func (c *PrometheusCollector) IncrementErrors() {
	promErrors.WithLabelValues(c.circuitName).Inc()
}

// IncrementSuccesses increments the number of requests that succeed.
func (c *PrometheusCollector) IncrementSuccesses() {
	promSuccesses.WithLabelValues(c.circuitName).Inc()
}

// IncrementFailures increments the number of requests that fail.
func (c *PrometheusCollector) IncrementFailures() {
	promFailures.WithLabelValues(c.circuitName).Inc()
}

// IncrementRejects increments the number of requests that are rejected.
func (c *PrometheusCollector) IncrementRejects() {
	promRejects.WithLabelValues(c.circuitName).Inc()
}

// IncrementShortCircuits increments the number of requests that short circuited
// due to the circuit being open.
func (c *PrometheusCollector) IncrementShortCircuits() {
	promShortCircuits.WithLabelValues(c.circuitName).Inc()
}

// IncrementTimeouts increments the number of timeouts that occurred in the
// circuit breaker.
func (c *PrometheusCollector) IncrementTimeouts() {
	promTimeouts.WithLabelValues(c.circuitName).Inc()
}

// IncrementFallbackSuccesses increments the number of successes that occurred
// during the execution of the fallback function.
func (c *PrometheusCollector) IncrementFallbackSuccesses() {
	promFallbackSuccesses.WithLabelValues(c.circuitName).Inc()
}

// IncrementFallbackFailures increments the number of failures that occurred
// during the execution of the fallback function.
func (c *PrometheusCollector) IncrementFallbackFailures() {
	promFallbackFailures.WithLabelValues(c.circuitName).Inc()
}

// UpdateTotalDuration updates the internal counter of how long we've run for.
func (c *PrometheusCollector) UpdateTotalDuration(timeSinceStart time.Duration) {
	promTotalDuration.WithLabelValues(c.circuitName).Set(timeSinceStart.Seconds())
}

// UpdateRunDuration updates the internal counter of how long the last run took.
func (c *PrometheusCollector) UpdateRunDuration(runDuration time.Duration) {
	promRunDuration.WithLabelValues(c.circuitName).Add(runDuration.Seconds())
}

// Reset is a noop operation in this collector.
func (c *PrometheusCollector) Reset() {}
