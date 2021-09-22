package main

import "fmt"

// 选择排序
func main() {
	arr := []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	arr = selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) []int {
	var result []int
	count := len(arr)

	for i := 0; i < count; i++ {
		smallestIndex := findSmallest(arr)
		result = append(result, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return result
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}
