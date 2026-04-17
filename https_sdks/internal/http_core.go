package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/rolandhe/go-base/commons"
	"github.com/rolandhe/go-base/https_sdks/http_define"
	"github.com/rolandhe/go-base/logger"
	"io"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

func Post[T, V any](hc *http.Client, bc *commons.BaseContext, req *T, url string, headers map[string]string, timeout time.Duration, level logger.LogLevel) (*V, error) {
	return do[*T, V](hc, bc, req, url, func() (*http.Request, context.CancelFunc, error) {
		var err error
		var r *http.Request
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer func() {
			if err != nil {
				cancel()
			}
		}()
		body, _ := json.Marshal(req)
		r, err = http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, nil, err
		}
		for k, v := range headers {
			r.Header[k] = []string{v}
		}
		r.Header["Content-Type"] = []string{"application/json"}
		return r, cancel, nil
	}, level)
}

func Get[V any](hc *http.Client, bc *commons.BaseContext, params map[string]string, headers map[string]string, url string, timeout time.Duration, level logger.LogLevel) (*V, error) {
	return do[map[string]string, V](hc, bc, params, url, func() (*http.Request, context.CancelFunc, error) {
		var err error
		var getRequest *http.Request
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer func() {
			if err != nil {
				cancel()
			}
		}()
		if len(params) > 0 {
			values := url2.Values{}
			for k, v := range params {
				values.Add(k, v)
			}
			if strings.Contains(url, "?") {
				url = url + "&" + values.Encode()
			} else {
				url = url + "?" + values.Encode()
			}
		}
		getRequest, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, nil, err
		}
		for k, v := range headers {
			getRequest.Header[k] = []string{v}
		}
		return getRequest, cancel, nil
	}, level)
}

func do[T, V any](hc *http.Client, bc *commons.BaseContext, param T, url string, reqFunc func() (*http.Request, context.CancelFunc, error), level logger.LogLevel) (*V, error) {
	start := time.Now().UnixMilli()
	if level&logger.LOG_LEVEL_PARAM == logger.LOG_LEVEL_PARAM {
		body, _ := json.Marshal(param)
		logger.WithBaseContextInfof(bc)("start to request %s, with body:%s", url, string(body))
	}

	var httpReq *http.Request
	var cancelFunc context.CancelFunc
	var err error
	var content []byte

	httpReq, cancelFunc, err = reqFunc()
	if err != nil {
		return nil, err
	}
	defer cancelFunc()

	defer func() {
		cost := time.Now().UnixMilli() - start
		if level&logger.LOG_LEVEL_RETURN == logger.LOG_LEVEL_RETURN {
			logger.WithBaseContextInfof(bc)("post end cost %d ms, with return:%s,err:%v", cost, string(content), err)
		} else if err != nil {
			logger.WithBaseContextInfof(bc)("post end cost %d ms, err:%v", cost, err)
		}
	}()

	var retValue *V

	retValue, content, err = CallWithResult[V](hc, bc, httpReq, nil)
	return retValue, err
}

func CallWithResult[V any](hc *http.Client, bc *commons.BaseContext, req *http.Request, noOKCallback http_define.ResponseNoOKCallback) (*V, []byte, error) {
	content, err := call(hc, bc, req, noOKCallback)
	if err != nil {
		return nil, nil, err
	}

	if len(content) == 0 {
		logger.WithBaseContextInfof(bc)("没有返回数据")
		return nil, nil, nil
	}

	var v = new(V)
	if shouldBytes(v, content) {
		return v, content, nil
	}
	if shouldString(v, content) {
		return v, content, nil
	}
	err = json.Unmarshal(content, v)
	return v, content, err
}

func shouldString(v any, content []byte) bool {
	p, ok := v.(*string)
	if ok {
		*p = string(content)
	}
	return ok
}

func shouldBytes(v any, content []byte) bool {
	p, ok := v.(*[]byte)
	if ok {
		*p = content
	}
	return ok
}

func call(hc *http.Client, bc *commons.BaseContext, req *http.Request, noOKCallback http_define.ResponseNoOKCallback) ([]byte, error) {
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	// ignore
	if resp.StatusCode == http.StatusNotFound {
		forNo200Content(bc, resp)
		return nil, http_define.Status404Err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		if noOKCallback != nil {
			processed, err := noOKCallback(resp)
			if err != nil {
				return nil, err
			}
			if processed {
				return nil, nil
			}
		}
		forNo200Content(bc, resp)
		return nil, errors.New(resp.Status)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func forNo200Content(bc *commons.BaseContext, resp *http.Response) {
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.WithBaseContextInfof(bc)("forNo200Content, status %d, to read content err:%v", resp.StatusCode, err)
		return
	}
	if len(content) == 0 {
		logger.WithBaseContextInfof(bc)("forNo200Content  , status %d, read empty content", resp.StatusCode)
		return
	}
	logger.WithBaseContextInfof(bc)("forNo200Content status:%d, content:%s", resp.StatusCode, string(content))
}
