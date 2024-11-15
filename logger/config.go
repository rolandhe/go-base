package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
)

var core *zap.SugaredLogger

var LogConfig = &struct {
	Path                string
	LogRotateMaxSize    int
	LogRotateMaxBackups int
	LogRotateMaxAge     int
	LogWithGid          bool
}{
	Path:                "stdout",
	LogRotateMaxSize:    500,
	LogRotateMaxBackups: 10,
	LogRotateMaxAge:     14,
	LogWithGid:          true,
}

func InitLogger() {
	var config zap.Config
	//logPath := ppath

	config = zap.NewDevelopmentConfig()
	config.OutputPaths = []string{LogConfig.Path}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	isStdout := LogConfig.Path == "stdout" || LogConfig.Path == "stderr"

	if isStdout {
		logger, err := config.Build()

		if err != nil {
			panic(err.Error())
		}
		core = logger.Sugar()
		return
	}

	info, err := os.Stat(LogConfig.Path)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err.Error())
		}
		if err = os.Mkdir(LogConfig.Path, os.ModePerm); err != nil {
			panic(err.Error())
		}
	} else if !info.IsDir() {
		panic("log path must be dir")
	}

	// 配置日志滚动
	lumberjackLogger := &lumberjack.Logger{
		Filename:   path.Join(LogConfig.Path, "out.log"), // 日志文件路径
		MaxSize:    LogConfig.LogRotateMaxSize,           // 每个日志文件的最大尺寸（MB）
		MaxBackups: LogConfig.LogRotateMaxBackups,        // 保留的旧日志文件的最大数量
		MaxAge:     LogConfig.LogRotateMaxAge,            // 保留旧日志文件的最大天数
		Compress:   false,                                // 是否压缩旧的日志文件
		LocalTime:  true,
	}

	// 自定义日志输出方式
	ccore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config.EncoderConfig),
		zapcore.AddSync(lumberjackLogger),
		config.Level,
	)
	core = zap.New(ccore, zap.AddCaller()).Sugar()
}

func Sync() {
	if core != nil {
		core.Sync()
	}
}
