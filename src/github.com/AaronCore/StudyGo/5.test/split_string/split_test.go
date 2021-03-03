package split_string

import (
	"reflect"
	"testing"
)

// 单元测试格式要求：
// 1.文件名必须以xx_test.go命名
// 2.方法必须是Test[^a-z]开头
// 3.方法参数必须 t *testing.T
// 4.使用go test执行单元测试
// 5.测试函数名必须以Test开头，必须接收一个*testing.T类型参数

// TestSplitString 单元测试
//func TestSplitString(t *testing.T) {
//	got := SplitString("a.b.c.d", ".")   // 程序输出的结果
//	want := []string{"a", "b", "c", "d"} // 期望的结果
//	if !reflect.DeepEqual(want, got) {   // 因为slice不能比较直接，借助反射包中的方法比较
//		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
//	}
//}

// TestSplitString 测试组
//func TestSplitString(t *testing.T) {
//	// 定义一个测试用例类型
//	type test struct {
//		input string
//		sep   string
//		want  []string
//	}
//	// 定义一个存储测试用例的切片
//	tests := []test{
//		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
//		{input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
//	}
//	// 遍历切片，逐一执行测试用例
//	for _, tc := range tests {
//		got := SplitString(tc.input, tc.sep)
//		if !reflect.DeepEqual(got, tc.want) {
//			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
//		}
//	}
//}

// TestSplitString 子测试
func TestSplitString(t *testing.T) {
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

// TestSplitString 基准测试
