package metrics

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics interface {
	IncHits(status int, method, path string)
	ObserveResponseTime(status int, method, path string, observeTime float64)
	IncRestPanicsTotal()
	IncGrpcPanicsTotal()
}

type EmptyMetrics struct{}

func (m *EmptyMetrics) IncHits(_ int, _, _ string)                        {}
func (m *EmptyMetrics) ObserveResponseTime(_ int, _, _ string, _ float64) {}
func (m *EmptyMetrics) IncRestPanicsTotal()                               {}
func (m *EmptyMetrics) IncGrpcPanicsTotal()                               {}

type PrometheusMetrics struct {
	HitsTotal             prometheus.Counter
	Hits                  *prometheus.CounterVec
	Times                 *prometheus.HistogramVec
	RestPanicRecoverTotal prometheus.Counter
	GrpcPanicRecoverTotal prometheus.Counter
}

func CreateMetrics(name string) (Metrics, error) {
	var metr PrometheusMetrics
	metr.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: name + "_hits_total",
	})
	if err := prometheus.Register(metr.HitsTotal); err != nil {
		return nil, err
	}

	metr.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name + "_hits",
		},
		[]string{"status", "method", "path"},
	)
	if err := prometheus.Register(metr.Hits); err != nil {
		return nil, err
	}

	metr.RestPanicRecoverTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: name + "_rest_panic_recover_total",
	})
	if err := prometheus.Register(metr.RestPanicRecoverTotal); err != nil {
		return nil, err
	}

	metr.GrpcPanicRecoverTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: name + "_grpc_panic_recover_total",
	})
	if err := prometheus.Register(metr.GrpcPanicRecoverTotal); err != nil {
		return nil, err
	}

	metr.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name + "_times",
		},
		[]string{"status", "method", "path"},
	)
	if err := prometheus.Register(metr.Times); err != nil {
		return nil, err
	}

	if err := prometheus.Register(collectors.NewBuildInfoCollector()); err != nil {
		return nil, err
	}

	return &metr, nil
}

type MetricsServerConfig struct {
	Host string `yaml:"host" env:"METRIC_HOST"`
	Port string `yaml:"port" env:"METRIC_PORT"`
}

func RunMetricServer(cfg MetricsServerConfig) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{Handler: http.Handler(mux), ReadHeaderTimeout: time.Second * 10}
	return srv.Serve(lis)
}

func (metr *PrometheusMetrics) IncHits(status int, method, path string) {
	metr.HitsTotal.Inc()
	metr.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

func (metr *PrometheusMetrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	metr.Times.WithLabelValues(strconv.Itoa(status), method, path).Observe(observeTime)
}

func (metr *PrometheusMetrics) IncRestPanicsTotal() {
	metr.RestPanicRecoverTotal.Inc()
}

func (metr *PrometheusMetrics) IncGrpcPanicsTotal() {
	metr.GrpcPanicRecoverTotal.Inc()
}
