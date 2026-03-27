package http2

import (
	"crypto/tls"
	"github.com/rolandhe/go-base/commons"
	"github.com/rolandhe/go-base/https_sdks/http_define"
	"github.com/rolandhe/go-base/https_sdks/internal"
	"github.com/rolandhe/go-base/logger"
	"golang.org/x/net/http2"
	"net/http"
	"sync"
	"time"
)

var http2Client *http.Client
var onceInit sync.Once

var limitConfig = struct {
	maxIdleConnsPerHost int
	maxIdleConns        int
}{5, 15}

func InitLimitConfig(maxIdleConnsPerHost, maxIdleConns int) {
	limitConfig.maxIdleConnsPerHost = maxIdleConnsPerHost
	limitConfig.maxIdleConns = maxIdleConns
}

func ensureHttpClient() {
	onceInit.Do(func() {
		if http2Client != nil {
			return
		}
		maxIdleConnsPerHost := limitConfig.maxIdleConnsPerHost
		maxIdleConns := limitConfig.maxIdleConns
		httpTransport := &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: false},
			IdleConnTimeout:     time.Minute * 30,
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
			MaxConnsPerHost:     maxIdleConns,
		}
		internal.ConfigureProxySupport(httpTransport)
		err := http2.ConfigureTransport(httpTransport)
		if err != nil {
			panic(err)
		}
		http2Client = &http.Client{
			Transport: httpTransport,
		}
	})
}

func GetHttpClient() *http.Client {
	ensureHttpClient()
	return http2Client
}

func Post[T, V any](bc *commons.BaseContext, req *T, url string, headers map[string]string, timeout time.Duration, level logger.LogLevel) (*V, error) {
	ensureHttpClient()
	return internal.Post[T, V](GetHttpClient(), bc, req, url, headers, timeout, level)
}

func Get[V any](bc *commons.BaseContext, params map[string]string, headers map[string]string, url string, timeout time.Duration, level logger.LogLevel) (*V, error) {
	ensureHttpClient()
	return internal.Get[V](GetHttpClient(), bc, params, headers, url, timeout, level)
}

func CallWithResult[V any](bc *commons.BaseContext, req *http.Request, noOKCallback http_define.ResponseNoOKCallback) (*V, []byte, error) {
	ensureHttpClient()
	return internal.CallWithResult[V](GetHttpClient(), bc, req, noOKCallback)
}
