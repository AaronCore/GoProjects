package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//writeFile()
	//bufioWriteFile()
	ioutilWriteFile()
}

func writeFile() {
	file, err := os.OpenFile("C:\\Log\\xxx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file error：", err)
		return
	}
	defer file.Close()
	content := "aaaaa\n"
	file.Write([]byte(content))
	file.WriteString("深圳市福田区")
}

func bufioWriteFile() {
	fileObj, err := os.OpenFile("C:\\Log\\xxx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file error：", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.Write([]byte("嘿嘿嘿...\n")) //将数据先写入缓存
	wr.WriteString("hello abc")
	wr.Flush() //将缓存中的内容写入文件
}

func ioutilWriteFile() {
	str := "SaaS大"
	err := ioutil.WriteFile("C:\\Log\\xxx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println("open file error：", err)
		return
	}
}
