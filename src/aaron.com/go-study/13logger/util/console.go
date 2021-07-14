package util

// 往终端打印log

import (
	"fmt"
	"time"
)

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLogger ConsoleLogger 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLoggerLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) log(level LogLevel, format string, a ...interface{}) {
	if level >= c.Level {
		msg := fmt.Sprintf(format, a...) // 格式化输出
		levelStr := getLoggerString(level)
		funcName, fileName, lineNo := getRunInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n", time.Now().Format("2006-01-02 15:04:05"), levelStr, fileName, funcName, lineNo, msg)
	}
}

// 判断是否需要记录日志
//func (c ConsoleLogger) enable(level LogLevel) bool {
//	return level >= c.Level
//}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Warn(format string, a ...interface{}) {
	c.log(WARN, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
