package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}

	fmt.Printf("%p\n", &a)
	fmt.Printf("%T\n", &a)
}
