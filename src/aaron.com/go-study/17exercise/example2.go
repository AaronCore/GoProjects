package main

import "fmt"

func main() {
	// 假设给定字符串 12345 使用穷举算法生成指定长度的字符串，如 假设生成长度为3的所有组合 111 112 113 114 115 121 122 123 124 125 131 132 133 。。。
	s := C("123456", 3)
	fmt.Println(s)
}

func C(s string, length int) []string {
	var sl int = len(s)
	if sl < length || length <= 0 {
		return nil
	}

	var M []string = make([]string, 0, sl)
	for _, v := range s {
		M = append(M, string(v))
	}

	var R []string = make([]string, sl, 1024)
	copy(R, M)

	for ; length-1 > 0; length-- {
		var l = len(R)
		for i := 0; i < l; i++ {
			for j := 1; j < sl; j++ {
				R = append(R, R[i]+M[j])
			}
			R[i] = R[i] + M[0]
		}
	}
	return R
}
