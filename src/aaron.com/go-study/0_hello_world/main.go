package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello ", os.Args[1])
	} else {
		fmt.Println("hello go！")
	}
	os.Exit(0)
}
