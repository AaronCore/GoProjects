package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP服务端程序的处理流程：
// 1.监听端口
// 2.接收客户端请求建立链接
// 3.创建goroutine处理链接

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() //关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("reader from client failed,err：", err)
			break
		}
		str := string(buf[:n])
		fmt.Println("收到client端发来的数据：", str)
		conn.Write([]byte(str)) // 发送数据
	}
}

func main() {
	//1.监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:200")
	if err != nil {
		fmt.Println("listen failed,err：", err)
		return
	}
	//2.接收客户端请求建立链接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err：", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
