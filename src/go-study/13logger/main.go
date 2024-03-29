package main

import (
	"go-study/13logger/util"
)

// 声明全局接口变量
var _ util.Logger

func main() {
	log := util.NewConsoleLogger("error")

	log.Debug("debug info...")
	log.Info("info...")
	log.Info("name：%v,age:%v", "aaron", 22)
	log.Error("error info...")
	log.Trace("trace info...")
	log.Warn("warn info...")
	log.Fatal("fatal info...")

	//time.Sleep(2 * time.Second)

	fileLog := util.NewFileLogger("info", "./", "dev.txt", 10*1024*1024)
	for i := 0; i < 100; i++ {
		fileLog.Info("深圳市福田区...")
	}
}
