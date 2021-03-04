package main

import "fmt"

// panic 和 recover，一般都是成对出现
// panic：抛出错误
//	1、内置函数
//  2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
//  3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
//  4、直到goroutine整个退出，并报告错误
// recover：捕获错误
//	1、内置函数
//	2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
//	3、一般的调用建议
//		a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
//      b). 可以获取通过panic传递的error
// 注意事项：
//	1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
//  2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
//  3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。
func main() {
	funcA()
	funcB()
	funcC()
}

func funcA() {
	fmt.Println("AA")
}
func funcB() {
	defer func() {
		err := recover()
		fmt.Println("err：", err)
		fmt.Println("资源释放...")
	}()
	panic("出现了严重的错误")
	fmt.Println("BB")
}
func funcC() {
	fmt.Println("CC")
}
