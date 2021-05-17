package logger_util

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

// Logger接口
type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Error(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

// 定义日志级别
const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	ERROR
	TRACE
	WARN
	FATAL
)

func parseLoggerLevel(levelStr string) (LogLevel, error) {
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "error":
		return ERROR, nil
	case "trace":
		return TRACE, nil
	case "warn":
		return WARN, nil
	case "fatal":
		return FATAL, nil
	default:
		return UNKNOWN, errors.New("invalid log level")
	}
}

func getLoggerString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case TRACE:
		return "TRACE"
	case WARN:
		return "WARN"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func getRunInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	return
}
