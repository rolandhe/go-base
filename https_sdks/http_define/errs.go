package http_define

import "errors"

var StatusCodeErr = errors.New("status code error")
var Status403Err = errors.New("status code 403")
var Status404Err = errors.New("status code 404")
