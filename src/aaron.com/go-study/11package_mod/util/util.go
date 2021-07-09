package util

import "fmt"

// 初始化
// 1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
// 2 每个包可以拥有多个init函数
// 3 包的每个源文件也可以拥有多个init函数
// 4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
// 5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
// 6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用

// init函数和main函数的异同：
// 相同点：
//   两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
// 不同点：
//   init可以应用于任意包中，且可以重复定义多个。
//   main函数只能用于main包中，且只能定义一个。
func init() {
	fmt.Println("初始调用util person")
}

// Sum
// 包中的标识符（变量名、函数名、结构体、接口...）如果首字母小写，表示私有（只能在当前包中使用）
// 首字母大写标识符，可以被外部使用
func Sum(x, y int) int {
	return x + y
}