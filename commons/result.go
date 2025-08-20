package commons

import (
	"errors"
)

type CodedResult interface {
	GetCode() int
}

type Result[T any] struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"errMsg"`
	Data   T      `json:"data,omitempty"`
}

func (r *Result[T]) GetCode() int {
	return r.Code
}

func QuickErrResult(errMsg string) *Result[*Void] {
	return &Result[*Void]{
		Code:   CommonErr,
		ErrMsg: errMsg,
	}
}

func ErrResult(code int, errMsg string) *Result[*Void] {
	return &Result[*Void]{
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

func QuickFromError(err error) *Result[*Void] {
	return FromError[*Void](err)
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
