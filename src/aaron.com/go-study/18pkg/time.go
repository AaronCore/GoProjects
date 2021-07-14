package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1、时间类型")
	now := time.Now()
	fmt.Println("当前时间：", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("时间：%d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)

	fmt.Println("2、时间戳")
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("时间戳：%v\n", timestamp1)
	fmt.Printf("纳秒时间戳：%v\n", timestamp2)

	fmt.Println("3、时间操作")
	thisTime := time.Now()
	fmt.Println("Add：", thisTime.Add(time.Hour)) // 给时间加时间
	fmt.Println("Sub：", thisTime.Sub(now))       // 计算时间差
	fmt.Println("Equal：", thisTime.Equal(now))   // 判断两个时间是否相同
	fmt.Println("Before：", thisTime.Before(now)) // 判断是否在时间点之前
	fmt.Println("After：", thisTime.After(now))   // 判断是否在时间点之后

	fmt.Println("4、定时器")
	//ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	//for i := range ticker {
	//	fmt.Println(i) //每秒都会执行的任务
	//}

	// 时间停留
	// 方式1
	time.Sleep(1 * time.Second)
	// 方式2
	a := 1000
	time.Sleep(time.Duration(a))

	fmt.Println("5、时间格式化")
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// time.Parse("2006-01-02 15:04:05", "2020-01-26 17:25:01")
	// 24小时制
	fmt.Println(now.Format("2006年01月02日 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	fmt.Println("5、练习，解析字符串格式的时间")
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
