package main

import "fmt"

func max(arr []int) int {
	maxValue := 0
	for _, item := range arr {
		if item > maxValue {
			maxValue = item
		}
	}
	fmt.Println(maxValue)
	return maxValue
}

func bs(arr []int) int {
	maxValue := max(arr)
	xs := make([]int, maxValue+1)
	for _, item := range arr {
		if xs[item] >= 1 {
			xs[item] = xs[item] + 1
		} else {
			xs[item] = 1
		}
	}

	for i, item := range xs {
		if xs[item] >= 1 {
			fmt.Println(i)
		}
	}
	return 0
}

func main() {
	bs([]int{1, 5, 3, 9, 7})
}
