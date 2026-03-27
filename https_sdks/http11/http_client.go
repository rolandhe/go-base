package http11

import (
	"crypto/tls"
	"github.com/rolandhe/go-base/commons"
	"github.com/rolandhe/go-base/https_sdks/http_define"
	"github.com/rolandhe/go-base/https_sdks/internal"
	"github.com/rolandhe/go-base/logger"
	"net/http"
	"sync"
	"time"
)

var httpClient *http.Client

var onceInit sync.Once

func ensureHttpClient() {
	onceInit.Do(func() {
		if httpClient != nil {
			return
		}
		httpTransport := &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: false},
			IdleConnTimeout:     time.Minute * 30,
			MaxIdleConnsPerHost: 5,
			MaxConnsPerHost:     15,
		}
		httpTransport.Proxy = nil
		httpClient = &http.Client{
			Transport: httpTransport,
		}
	})
}

func GetHttpClient() *http.Client {
	ensureHttpClient()
	return httpClient
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
