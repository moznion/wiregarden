package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PrometheusMetricsRegisterable interface {
	IncrementRequestCount(requestName string)
	IncrementSuccessResponseCount(requestName string)
	IncrementFailureResponseCount(requestName string, code uint32)
}

type PrometheusMetricsRegister struct {
	requestCountCounterOpts         prometheus.CounterOpts
	successResponseCountCounterOpts prometheus.CounterOpts
	failureResponseCountCounterOpts prometheus.CounterOpts
	requestCountCounterMap          map[string]prometheus.Counter
	successResponseCountCounterMap  map[string]prometheus.Counter
	failureResponseCountCounterMap  map[string]map[uint32]prometheus.Counter
}

func NewPrometheusMetricsRegister() *PrometheusMetricsRegister {
	return &PrometheusMetricsRegister{
		requestCountCounterOpts: prometheus.CounterOpts{
			Namespace: "wiregarden",
			Subsystem: "grpc",
			Name:      "request_count",
			Help:      "The number of gRPC requests",
		},
		successResponseCountCounterOpts: prometheus.CounterOpts{
			Namespace: "wiregarden",
			Subsystem: "grpc",
			Name:      "success_response_count",
			Help:      "The number of gRPC responses that are succeeded",
		},
		failureResponseCountCounterOpts: prometheus.CounterOpts{
			Namespace: "wiregarden",
			Subsystem: "grpc",
			Name:      "failure_response_count",
			Help:      "The number of gRPC responses that are failed",
		},
		requestCountCounterMap:         make(map[string]prometheus.Counter),
		successResponseCountCounterMap: make(map[string]prometheus.Counter),
		failureResponseCountCounterMap: make(map[string]map[uint32]prometheus.Counter),
	}
}

func (p *PrometheusMetricsRegister) IncrementRequestCount(requestName string) {
	counter, ok := p.requestCountCounterMap[requestName]
	if !ok {
		counter = promauto.NewCounter(prometheus.CounterOpts{
			Namespace: p.requestCountCounterOpts.Namespace,
			Subsystem: p.requestCountCounterOpts.Subsystem,
			Name:      p.requestCountCounterOpts.Name,
			Help:      p.requestCountCounterOpts.Help,
			ConstLabels: prometheus.Labels{
				"request_name": requestName,
			},
		})
		p.requestCountCounterMap[requestName] = counter
	}
	counter.Inc()
}

func (p *PrometheusMetricsRegister) IncrementSuccessResponseCount(requestName string) {
	counter, ok := p.successResponseCountCounterMap[requestName]
	if !ok {
		counter = promauto.NewCounter(prometheus.CounterOpts{
			Namespace: p.successResponseCountCounterOpts.Namespace,
			Subsystem: p.successResponseCountCounterOpts.Subsystem,
			Name:      p.successResponseCountCounterOpts.Name,
			Help:      p.successResponseCountCounterOpts.Help,
			ConstLabels: prometheus.Labels{
				"request_name": requestName,
			},
		})
		p.successResponseCountCounterMap[requestName] = counter
	}
	counter.Inc()
}

func (p *PrometheusMetricsRegister) IncrementFailureResponseCount(requestName string, code uint32) {
	counterMap, ok := p.failureResponseCountCounterMap[requestName]
	if !ok {
		counterMap = make(map[uint32]prometheus.Counter)
		p.failureResponseCountCounterMap[requestName] = counterMap
	}

	counter, ok := counterMap[code]
	if !ok {
		counter = promauto.NewCounter(prometheus.CounterOpts{
			Namespace: p.failureResponseCountCounterOpts.Namespace,
			Subsystem: p.failureResponseCountCounterOpts.Subsystem,
			Name:      p.failureResponseCountCounterOpts.Name,
			Help:      p.failureResponseCountCounterOpts.Help,
			ConstLabels: prometheus.Labels{
				"request_name": requestName,
				"code":         fmt.Sprintf("%d", code),
			},
		})
		counterMap[code] = counter
	}
	counter.Inc()
}

type NOPPrometheusMetricsRegister struct {
}

func (n *NOPPrometheusMetricsRegister) IncrementRequestCount(requestName string) {
}

func (n *NOPPrometheusMetricsRegister) IncrementSuccessResponseCount(requestName string) {
}

func (n *NOPPrometheusMetricsRegister) IncrementFailureResponseCount(requestname string, code uint32) {
}
