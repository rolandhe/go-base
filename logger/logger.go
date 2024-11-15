package logger

import (
	"bytes"
	"fmt"
	"github.com/rolandhe/go-base/commons"
	"go.uber.org/zap"
	"runtime"
	"strconv"
)

func getGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func WithBaseContextDebugf(bc commons.BaseContext) func(template string, args ...interface{}) {
	traceId := bc.Get(commons.TraceId)
	return func(template string, args ...interface{}) {
		template = changeTemplate(template, traceId)
		core.WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
	}
}

func Debugf(template string, args ...interface{}) {
	template = changeTemplate(template, "")
	core.WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
}

func WithBaseContextInfof(bc commons.BaseContext) func(template string, args ...interface{}) {
	traceId := bc.Get(commons.TraceId)
	return func(template string, args ...interface{}) {
		template = changeTemplate(template, traceId)
		core.WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
	}
}

func Infof(template string, args ...interface{}) {
	template = changeTemplate(template, "")
	core.WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
}

func WithBaseContextWarnf(bc commons.BaseContext) func(template string, args ...interface{}) {
	traceId := bc.Get(commons.TraceId)
	return func(template string, args ...interface{}) {
		template = changeTemplate(template, traceId)
		core.WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
	}
}
func Warnf(template string, args ...interface{}) {
	template = changeTemplate(template, "")
	core.WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
}

func WithBaseContextErrorf(bc commons.BaseContext) func(template string, args ...interface{}) {
	traceId := bc.Get(commons.TraceId)
	return func(template string, args ...interface{}) {
		template = changeTemplate(template, traceId)
		core.WithOptions(zap.AddCallerSkip(1)).Errorf(template, args...)
	}
}
func Errorf(template string, args ...interface{}) {
	template = changeTemplate(template, "")
	core.WithOptions(zap.AddCallerSkip(1)).Errorf(template, args...)
}

func changeTemplate(template string, traceId string) string {
	if traceId == "" {
		if LogConfig.LogWithGid {
			template = fmt.Sprintf("gid=%d,%s", getGoroutineID(), template)
		}
		return template
	}
	if LogConfig.LogWithGid {
		return fmt.Sprintf("tid=%s,gid=%d,%s", traceId, getGoroutineID(), template)
	}
	return fmt.Sprintf("tid=%s,%s", traceId, template)
}
