package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dirName := "C:\\Temp"
	listFiles(dirName, 0)
}

func listFiles(dirName string, level int) {
	// level用来记录当前递归的层次
	// 生成有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|   " + s
	}

	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirName + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			listFiles(filename, level+1) // 继续遍历fi这个目录
		}
	}
}
