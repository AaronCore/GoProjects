package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//readFile()
	//bufioReadFile()
	ioutilReadFile()
}

// 打开文件
func readFile() {
	file, err := os.Open("F:\\GoProjects\\src\\github.com\\AaronCore\\StudyGo\\1base\\1.string\\main.go")
	if err != nil {
		fmt.Println("open file error：", err)
		return
	}
	defer file.Close()
	var tmp [128]byte
	for {
		n, err := file.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("read end...")
			return
		}
		if err != nil {
			fmt.Println("read file error：", err)
			return
		}
		fmt.Printf("读取了%d字节数据\n", n)
		fmt.Printf(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// bufio打开文件
func bufioReadFile() {
	file, err := os.Open("F:\\GoProjects\\src\\github.com\\AaronCore\\StudyGo\\1base\\1.string\\main.go")
	if err != nil {
		fmt.Println("open file error：", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("read end...")
			return
		}
		if err != nil {
			fmt.Println("read file error：", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutilReadFile 打开文件
func ioutilReadFile() {
	content, err := ioutil.ReadFile("F:\\GoProjects\\src\\github.com\\AaronCore\\StudyGo\\1base\\1.string\\main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}
