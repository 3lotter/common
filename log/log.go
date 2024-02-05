// Package log
//
//	@Title			log.go
//	@Description	本文件用于初始化日志
package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"time"
)

type MyLogger struct {
	*logrus.Logger
}

// Conf 日志配置
type Conf struct {
	Level      string
	Path       string
	RotateTime int64
}

// SetupLogger 初始化日志
func SetupLogger(logConf Conf, formatter logrus.Formatter) (*MyLogger, error) {
	var err error

	// 日志输出格式
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)

	// 日志输出等级
	var level logrus.Level
	level, err = logrus.ParseLevel(logConf.Level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)

	// 日志输出位置
	if len(logConf.Path) == 0 {
		logger.SetOutput(os.Stdout)
	} else {
		var writer *rotatelogs.RotateLogs
		path := logConf.Path
		writer, err = rotatelogs.New(
			path+".%Y%m%d",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithRotationTime(time.Duration(logConf.RotateTime)*time.Hour),
		)
		if err != nil {
			return nil, err
		}
		logger.SetOutput(writer)
	}

	return &MyLogger{Logger: logger}, nil
}

// getCallerLocation 获取调用者的调用行
func (logger *MyLogger) getCallerLocation() string {
	// 调用链往上翻三层，找到调用函数的信息
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return ""
	}
	fileName := path.Base(file)
	funcName := runtime.FuncForPC(pc).Name()

	location := fmt.Sprintf("%s %s:%d", funcName, fileName, line)
	return location
}

// printerF 日志输出，记录报错的位置
func (logger *MyLogger) printerF(level logrus.Level, format string, args ...any) {
	location := logger.getCallerLocation()
	entry := logger.WithFields(logrus.Fields{})
	entry.Data[dataLocation] = location
	entry.Logf(level, format, args...)
}

// DebugF DEBUG级别日志输出
func (logger *MyLogger) DebugF(format string, args ...any) {
	logger.printerF(logrus.DebugLevel, format, args...)
}

// InfoF INFO级别日志输出
func (logger *MyLogger) InfoF(format string, args ...any) {
	logger.printerF(logrus.InfoLevel, format, args...)
}

// WarnF WARN级别日志输出
func (logger *MyLogger) WarnF(format string, args ...any) {
	logger.printerF(logrus.WarnLevel, format, args...)
}

// ErrorF ERROR级别日志输出
func (logger *MyLogger) ErrorF(format string, args ...any) {
	logger.printerF(logrus.ErrorLevel, format, args...)
}

// FatalF FATAL级别日志输出
func (logger *MyLogger) FatalF(format string, args ...any) {
	logger.printerF(logrus.FatalLevel, format, args...)
}
