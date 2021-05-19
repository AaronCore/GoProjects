package main

import (
	"aaron.com/go-study/13_logger/logger_util"
)

// 声明全局接口变量
var log logger_util.Logger

func main() {
	log := logger_util.NewConsoleLogger("error")

	log.Debug("debug info...")
	log.Info("info...")
	log.Info("name：%v,age:%v", "aaron", 22)
	log.Error("error info...")
	log.Trace("trace info...")
	log.Warn("warn info...")
	log.Fatal("fatal info...")

	//time.Sleep(2 * time.Second)

	fileLog := logger_util.NewFileLogger("info", "./", "dev.txt", 10*1024*1024)
	for i := 0; i < 100; i++ {
		fileLog.Info("深圳市福田区...")
	}
}
