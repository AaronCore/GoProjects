package main

import "fmt"

// 快速排序
func main() {
	arr := []int{12, 87, 1, 66, 30, 126, 128, 12, 653, 78, 98}
	result := quickSort(arr)
	fmt.Println(result)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, ele := range arr[1:] {
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}
