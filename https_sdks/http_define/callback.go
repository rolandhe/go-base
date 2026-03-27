package http_define

import "net/http"

type ResponseNoOKCallback func(resp *http.Response) (bool, error)
