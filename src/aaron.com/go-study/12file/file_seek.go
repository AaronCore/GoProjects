package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	/*
		断点续传：
			文件传递：文件复制
				C:\Temp\a\a.png
			复制到
				C:\Temp\b\b.png
		思路：
			边复制，边记录复制的总量
	*/

	srcFile := "C:\\Temp\\a\\a.png"
	destFile := "C:\\Temp\\b\\b.png"
	tempFile := destFile + "temp.txt"

	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)

	defer file1.Close()
	defer file2.Close()

	//1.读取临时文件中的数据，根据seek
	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 128, 128)
	n1, err := file3.Read(bs)
	fmt.Println("ni：", n1)

	countStr := string(bs[:n1])
	fmt.Println("countStr：", countStr)
	//count,_:=strconv.Atoi(countStr)
	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println("count：", count)

	//2. 设置读，写的偏移量
	file1.Seek(count, 0)
	file2.Seek(count, 0)
	data := make([]byte, 1024, 1024)
	n2 := -1            // 读取的数据量
	n3 := -1            //写出的数据量
	total := int(count) //读取的总量

	for {
		//3.读取数据
		n2, err = file1.Read(data)
		if err == io.EOF {
			fmt.Println("文件复制完毕。。")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		//将数据写入到目标文件
		n3, _ = file2.Write(data[:n2])
		total += n3
		//将复制总量，存储到临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		// 模拟失败
		//if total > 8000 {
		//	panic("取消上传了...")
		//}
	}
}
