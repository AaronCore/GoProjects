package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//writeFile()
	//buffWriteFile()
	utilWriteFile()
}

// writeFile
func writeFile() {
	file, err := os.OpenFile("C:\\Temp\\xxx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file error：", err)
		return
	}
	defer file.Close()
	content := "away\n"
	file.Write([]byte(content))
	file.WriteString("深圳市福田区")
}

// buffWriteFile
func buffWriteFile() {
	fileObj, err := os.OpenFile("C:\\Temp\\xxx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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

// utilWriteFile
func utilWriteFile() {
	str := "hello go go.."
	err := ioutil.WriteFile("C:\\Temp\\xxx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("open file error：", err)
		return
	}
}
