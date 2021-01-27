package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	// 配置日志输入位置
	logFile, err := os.OpenFile("C:\\Log\\log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed err：", err)
		return
	}
	// defer logFile.Close()
	log.SetOutput(logFile)
	// 设置打印格式
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")

	// 配置日志前缀
	log.SetPrefix("[log]")
	log.Println("这是条普通的日志...")
	log.Printf("这是一条%s日志。\n", "很普通的")
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
