package metrics

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//var (
//	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
//		Name: "cpu_temperature_celsius",
//		Help: "Current temperature of the CPU.",
//	})
//
//	counter = prometheus.NewCounter(prometheus.CounterOpts{
//		Name: "api_http_requests_total",
//		Help: "count http request",
//	})
//)

//func init() {
//	// Metrics have to be registered to be exposed:
//	prometheus.MustRegister(cpuTemp)
//	prometheus.MustRegister(counter)
//}

func MetricHandler() echo.HandlerFunc {

	return echo.WrapHandler(promhttp.Handler())

}
