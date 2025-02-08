package metrics

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var (
	p = fasthttpadaptor.NewFastHTTPHandler(
		promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		),
	)

	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of http requests.",
		},
		[]string{"status", "method", "path"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_secs",
			Help: "http request duration in seconds.",
		},
		[]string{"status", "method", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
}

func PrometheusHandler(c *fiber.Ctx) error {
	p(c.Context())
	return nil
}

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()
		method := c.Method()
		status := c.Context().Response.StatusCode()

		start := time.Now()
		duration := time.Since(start).Seconds()

		requestCount.WithLabelValues(strconv.Itoa(status), method, path).Inc()
		requestDuration.WithLabelValues(strconv.Itoa(status), method, path).Observe(duration)

		return c.Next()
	}
}
