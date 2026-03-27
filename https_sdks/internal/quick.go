package internal

import (
	"github.com/rolandhe/go-base/commons"
	"github.com/rolandhe/go-base/https_sdks/http_define"
	"io"
	"log"
	"net/http"
	"os"
)

func GetStringBodyResponse(hc *http.Client, bc *commons.BaseContext, req *http.Request, respCookieCallback func(cookies []*http.Cookie)) (string, error) {
	resp, err := hc.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		forNo200Content(bc, resp)
		log.Printf("http status code %d, %s", resp.StatusCode, req.URL.String())
		if resp.StatusCode == 403 {
			return "", http_define.Status403Err
		}
		if resp.StatusCode == 404 {
			return "", http_define.Status404Err
		}
		return "", http_define.StatusCodeErr
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if respCookieCallback != nil {
		respCookieCallback(resp.Cookies())
	}

	return string(body), nil
}

func DownloadFileMemory(hc *http.Client, bc *commons.BaseContext, req *http.Request) ([]byte, error) {
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		forNo200Content(bc, resp)
		if resp.StatusCode == 403 {
			return nil, http_define.Status403Err
		}
		if resp.StatusCode == 404 {
			return nil, http_define.Status404Err
		}
		return nil, http_define.StatusCodeErr
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DownloadFile(hc *http.Client, bc *commons.BaseContext, req *http.Request, outFile string) error {
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		//log.Printf("http status code %d, %s", resp.StatusCode, req.URL.String())
		forNo200Content(bc, resp)
		if resp.StatusCode == 403 {
			return http_define.Status403Err
		}
		if resp.StatusCode == 404 {
			return http_define.Status404Err
		}
		return http_define.StatusCodeErr
	}

	var w *os.File
	w, err = os.Create(outFile)
	if err != nil {
		return err
	}
	var writeErr error
	defer func() {
		if err = w.Close(); err != nil {
			log.Printf("close file %s, err: %v", outFile, err)
			return
		}
		if writeErr != nil {
			_ = os.Remove(outFile)
		}
	}()
	_, writeErr = io.Copy(w, resp.Body)

	return writeErr
}
