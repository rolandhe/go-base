package commons

import "go/types"

const (
	OKCode    int = 200
	NotLogin      = 4001
	CommonErr     = 5000
)

type Result[T any] struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"errMsg"`
	Data   T      `json:"data"`
}

func QuickErrResult(errMsg string) *Result[types.Nil] {
	return &Result[types.Nil]{
		Code:   CommonErr,
		ErrMsg: errMsg,
	}
}

func ErrResult(code int, errMsg string) *Result[types.Nil] {
	return &Result[types.Nil]{
		Code:   code,
		ErrMsg: errMsg,
	}
}

func OkResult[T any](data T) *Result[T] {
	return &Result[T]{
		Code: OKCode,
		Data: data,
	}
}

func NewResult[T any](code int, errMsg string, data T) *Result[T] {
	return &Result[T]{
		Code:   code,
		ErrMsg: errMsg,
		Data:   data,
	}
}
