package monitor

import (
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rolandhe/go-base/logger"
	"net/http"
	"sync"
)

var ServerReqCounter *prometheus.CounterVec
var ServerReqGauge *prometheus.GaugeVec
var ServerReqDuration *prometheus.HistogramVec

var ClientReqCounter *prometheus.CounterVec
var ClientReqGauge *prometheus.GaugeVec
var ClientReqDuration *prometheus.HistogramVec

var BizEventCounter *prometheus.CounterVec

var counterMap = new(sync.Map)

func StartMonitor(appName string, port int) {
	ServerReqCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "server_req_count",
		Help: "server side request counter",
		ConstLabels: map[string]string{
			"appName": appName,
		},
	}, []string{"path"})
	ServerReqGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "server_response_duration_seconds",
		Help: "Duration to server requests.",
	}, []string{"path", "error"})
	ServerReqDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "server_response_duration_seconds_histogram",
		Help: "Duration to server requests.",
	}, []string{"path", "error"})

	ClientReqCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "client_req_count",
		Help: "client request counter",
	}, []string{"path"})
	ClientReqGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "client_response_duration_seconds",
		Help: "Duration to client requests.",
	}, []string{"path", "error"})
	ClientReqDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "client_response_duration_seconds_histogram",
		Help: "Duration to client requests.",
	}, []string{"path", "error"})

	BizEventCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "biz_event_count",
		Help: "biz side event counter",
	}, []string{"event", "error"})
	ListenAndServe(port)
}

func ListenAndServe(port int) {
	go func() {
		http.Handle("/monitor/prometheus", promhttp.Handler())
		host := fmt.Sprintf(":%d", port)
		logger.Infof("start to monitor:%d....", port)
		if err := http.ListenAndServe(host, nil); err != nil {
			logger.Errorf("monitor listen failed, port:%d, err:%v", port, err)
		}
	}()
}

func DoClientReqCounter(url string) {
	if ClientReqCounter == nil {
		return
	}
	ClientReqCounter.WithLabelValues(url).Inc()
}

func DoClientDuration(url string, e string, cost int64) {
	if ClientReqDuration == nil || ClientReqGauge == nil {
		return
	}
	ClientReqDuration.WithLabelValues(url, e).Observe(float64(cost))
	ClientReqGauge.WithLabelValues(url, e).Set(float64(cost))
}

func DoServerCounter(url string) {
	if ServerReqCounter == nil {
		return
	}
	ServerReqCounter.WithLabelValues(url).Inc()
}

func DoBizEventCounter(event string, e string) {
	if BizEventCounter == nil {
		return
	}
	BizEventCounter.WithLabelValues(event, e).Inc()
}

func DoServerDuration(url string, e string, cost int64) {
	if ServerReqDuration == nil || ServerReqGauge == nil {
		return
	}
	ServerReqDuration.WithLabelValues(url, e).Observe(float64(cost))
	ServerReqGauge.WithLabelValues(url, e).Set(float64(cost))
}

func IncCounter(name string, labelMap map[string]string) error {
	if name == "" || len(labelMap) == 0 {
		return errors.New("invalid params")
	}

	labelKeys := make([]string, 0, len(labelMap))
	labelValues := make([]string, 0, len(labelMap))
	for k, v := range labelMap {
		labelKeys = append(labelKeys, k)
		labelValues = append(labelValues, v)
	}

	actual, _ := counterMap.LoadOrStore(name, promauto.NewCounterVec(prometheus.CounterOpts{
		Name: name,
		Help: name,
	}, labelKeys))

	actual.(*prometheus.CounterVec).WithLabelValues(labelValues...).Inc()
	return nil
}
