package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("1、string--->int")
	fmt.Println(strconv.Atoi("10"))

	fmt.Println("2、int--->string")
	fmt.Println(strconv.Itoa(20))

	fmt.Println("3、Parse系列函数")
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseInt("10", 10, 64))
	fmt.Println(strconv.ParseFloat("3.14", 64))
	fmt.Println(strconv.ParseUint("3", 10, 64))

	fmt.Println("4、Format系列函数")
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatFloat(3.1415, 'E', -1, 64))
	fmt.Println(strconv.FormatInt(-2, 16))
	fmt.Println(strconv.FormatUint(2, 16))

	fmt.Println("5、其它函数")
	fmt.Println(strconv.IsPrint('\u263a'))
}
