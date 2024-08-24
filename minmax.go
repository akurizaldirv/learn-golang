package main

import "fmt"

func getMin(arr []int) int {
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}

func getMax(arr []int) int {
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func main() {
	arr := []int{12, 34, 23, 4, 56, 25}
	fmt.Println(getMin(arr))
	fmt.Println(getMax(arr))
}
