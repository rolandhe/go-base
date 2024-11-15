package logger

import (
	"bytes"
	"fmt"
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

func WithTraceIdDebugf(traceId string) func(template string, args ...interface{}) {
	return func(template string, args ...interface{}) {
		template = changeTemplate(template, traceId)
		core.WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
	}
}

func Debugf(template string, args ...interface{}) {
	template = changeTemplate(template, "")
	core.WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
}

func WithTraceIdInfof(traceId string) func(template string, args ...interface{}) {
	return func(template string, args ...interface{}) {
		template = changeTemplate(template, traceId)
		core.WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
	}
}

func Infof(template string, args ...interface{}) {
	template = changeTemplate(template, "")
	core.WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
}

func WithTraceIdWarnf(traceId string) func(template string, args ...interface{}) {
	return func(template string, args ...interface{}) {
		template = changeTemplate(template, traceId)
		core.WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
	}
}
func Warnf(template string, args ...interface{}) {
	template = changeTemplate(template, "")
	core.WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
}

func WithTraceIdErrorf(traceId string) func(template string, args ...interface{}) {
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
