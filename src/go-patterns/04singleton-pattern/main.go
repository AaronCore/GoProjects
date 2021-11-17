package main

import "fmt"

// 单例模式
func main() {
	for i := 0; i < 100; i++ {
		go GetInstance()
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
