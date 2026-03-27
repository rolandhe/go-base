package http11

import (
	"github.com/rolandhe/go-base/commons"
	"github.com/rolandhe/go-base/https_sdks/internal"
	"net/http"
)

func GetStringBodyResponse(bc *commons.BaseContext, req *http.Request, respCookieCallback func(cookies []*http.Cookie)) (string, error) {
	ensureHttpClient()
	return internal.GetStringBodyResponse(httpClient, bc, req, respCookieCallback)
}

func DownloadFileMemory(bc *commons.BaseContext, req *http.Request) ([]byte, error) {
	ensureHttpClient()
	return internal.DownloadFileMemory(httpClient, bc, req)
}

func DownloadFile(bc *commons.BaseContext, req *http.Request, outFile string) error {
	ensureHttpClient()
	return internal.DownloadFile(httpClient, bc, req, outFile)
}
