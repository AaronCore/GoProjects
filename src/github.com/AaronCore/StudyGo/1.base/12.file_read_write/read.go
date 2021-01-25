package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//readFromFile()
	//readFromFileBufio()
	readFromFileIoutil()
}

// 打开文件
func readFromFile() {
	file, err := os.Open("F:\\GoProjects\\src\\github.com\\AaronCore\\StudyGo\\1.base\\1.string\\main.go")
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
func readFromFileBufio() {
	file, err := os.Open("F:\\GoProjects\\src\\github.com\\AaronCore\\StudyGo\\1.base\\1.string\\main.go")
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

// ioutil打开文件
func readFromFileIoutil() {
	content, err := ioutil.ReadFile("F:\\GoProjects\\src\\github.com\\AaronCore\\StudyGo\\1.base\\1.string\\main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}