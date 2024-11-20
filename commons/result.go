package commons

import (
	"errors"
	"go/types"
)

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

func QuickTypeErrResult[T any](errMsg string) *Result[T] {
	return &Result[T]{
		Code:   CommonErr,
		ErrMsg: errMsg,
	}
}

func ErrTypeResult[T any](code int, errMsg string) *Result[T] {
	return &Result[T]{
		Code:   code,
		ErrMsg: errMsg,
	}
}

func FromError[T any](err error) *Result[T] {
	var stdErr *StdError
	ok := errors.As(err, &stdErr)
	if ok {
		return &Result[T]{
			Code:   stdErr.Code,
			ErrMsg: stdErr.Message,
		}
	}
	return &Result[T]{
		Code:   CommonErr,
		ErrMsg: "internal server error",
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
