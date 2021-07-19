package split_string

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// 单元测试格式要求：
// 1.文件名必须以xx_test.go命名
// 2.方法必须是Test[^a-z]开头
// 3.方法参数必须 t *testing.T
// 4.使用go test执行单元测试
// 5.测试函数名必须以Test开头，必须接收一个*testing.T类型参数

// 测试覆盖率
// 1. 查看测试覆盖率：go test -cover
// 2. 将测试覆盖率输出到文件：go test -cover -coverprofile=c.out

// TestSplitString 单元测试
func TestSplitString(t *testing.T) {
	got := SplitString("a.b.c.d", ".")   // 程序输出的结果
	want := []string{"a", "b", "c", "d"} // 期望的结果
	if !reflect.DeepEqual(want, got) {   // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

// TestSplitString1 测试组
func TestSplitString1(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := SplitString(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
		}
	}
}

// TestSplitString2 子测试
func TestSplitString2(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"case1": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"case2": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"case3": {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"case4": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := SplitString(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 基准测试
// go test -bench=SplitString -benchmem
// TestSplitString 基准测试
func BenchmarkSplitString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitString("a.b.c.d.e.f", ".")
	}
}

// 重置时间
func BenchmarkSplitString2(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		SplitString("a.b.c.d.e.f", ".")
	}
}

// 并行测试
func BenchmarkSplitString3(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SplitString("a.b.c.d.e.f", ".")
		}
	})
}

// 性能比较
// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}
func BenchmarkFib1(b *testing.B) { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B) { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B) { benchmarkFib(b, 3) }

// 示例函数
// 示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。
// go test -run Example
// ExampleSplitString
func ExampleSplitString() {
	fmt.Println(SplitString("a:b:c", ":"))
}
