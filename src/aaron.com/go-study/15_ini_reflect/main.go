package main

import (
	"aaron.com/go-study/15_ini_reflect/util"
	"fmt"
	"io/ioutil"
)

// 解析文件
func parseFile1(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	var conf util.Config
	err = util.UnMarshal(data, &conf)
	if err != nil {
		return
	}
	fmt.Printf("反序列化成功  conf: %#v\n  port: %#v\n", conf, conf.SqlServerConf.Port)

}

func parseFile2(fileName string) {
	conf := util.Config{
		SqlServerConf: util.SqlServerConfig{
			Ip:   "127.0.0.1",
			Port: 8500,
		},
		MySqlConf: util.MySqlConfig{
			UserName: "127.0.0.1",
			PassWord: "123456",
			DataBase: "dev",
			Host:     "127.0.0.1",
			Port:     3306,
			TimeOut:  3600,
		},
	}
	err := util.MarshalFile(fileName, conf)
	if err != nil {
		return
	}
}

// 配置文件反射
func main() {
	parseFile1("F:\\GoProjects\\src\\github.com\\AaronCore\\StudyGo\\1base\\15ini_reflect\\config.ini")
	//parseFile2("D:\\my.ini")
}
