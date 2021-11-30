package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

// doc：https://ini.unknwon.cn/
func main() {
	cnf := new(AppConfig)

	// 方式1
	//cfg, err := ini.Load("config.ini")
	//if err != nil {
	//	fmt.Printf("read file failed,err:%v", err)
	//	os.Exit(1)
	//}
	//_ = cfg.MapTo(cnf)

	// 方式2
	_ = ini.MapTo(cnf, "config.ini")

	fmt.Println(cnf.MySQLConfig)
}

// baseExample 基本示例操作
func baseExample() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("read file failed,err:%v", err)
		os.Exit(1)
	}

	// 一般读取
	fmt.Println("App Mode：", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path：", cfg.Section("paths").Key("data").String())

	// 候选值限制的操作
	fmt.Println("Server Protocol:", cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// 如果读取的值不在候选列表内，则会回退使用提供的默认值
	fmt.Println("Email Protocol:", cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// 自动类型转换
	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// 修改某个值然后进行保存
	cfg.Section("").Key("app_mode").SetValue("production")
	_ = cfg.SaveTo("my.ini.local")
}
