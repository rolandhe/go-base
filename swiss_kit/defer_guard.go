package swiss_kit

import (
	"errors"
	"fmt"
	"github.com/rolandhe/go-base/commons"
	"github.com/rolandhe/go-base/logger"
	"runtime"
)

func DeferGuardFuncWithSlice[T any](bc *commons.BaseContext, coreFunc func() (res []T, err error)) (res []T, err error) {
	defer func() {
		if r := recover(); r != nil {
			stack := CaptureStack(3)
			logger.WithBaseContextInfof(bc)("met panic err:%v,stack:%s", r, stack)
			err = errors.New("panic err")
			res = nil
		}
	}()
	return coreFunc()
}

func DeferGuardFunc[T any](bc *commons.BaseContext, coreFunc func() (res *T, err error)) (res *T, err error) {
	defer func() {
		if r := recover(); r != nil {
			stack := CaptureStack(3)
			logger.WithBaseContextInfof(bc)("met panic err:%v,stack:\n%s", r, stack)
			err = errors.New("panic err")
			res = nil
		}
	}()
	return coreFunc()
}

// CaptureStack 获取当前 goroutine 的堆栈信息
// skip: 跳过前 skip 层调用，一般 skip=3-4 可以去掉 runtime.Callers + defer 层
func CaptureStack(skip int) string {
	const maxDepth = 64
	var pcs [maxDepth]uintptr

	// 获取调用栈
	n := runtime.Callers(skip, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	stack := ""
	for {
		frame, more := frames.Next()
		stack += fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}

	return stack
}
