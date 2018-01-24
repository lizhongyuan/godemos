package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	var numbers2 []int
	if numbers2 == nil {
		fmt.Printf("the slice is nil")
	}
}

func printSlice(x []int) {
	fmt.Printf("len:%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
