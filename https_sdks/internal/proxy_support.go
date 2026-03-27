package internal

import (
	"net/http"
	"os"
	"strings"
)

const (
	useHttpProxy = "USE_HTTP_PROXY"
)

func ConfigureProxySupport(transport *http.Transport) {
	useProxy := os.Getenv(useHttpProxy)
	if useProxy == "" {
		useProxy = os.Getenv(strings.ToLower(useHttpProxy))
	}
	if useProxy == "true" {
		return
	}
	transport.Proxy = nil
}
