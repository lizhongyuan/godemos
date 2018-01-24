package main

import "fmt"


func main() {
	numbers := []int {0, 1, 2, 3}

	numbers2 := numbers[:2]

	printSlice2(numbers)

	fmt.Print(numbers[:1])
	fmt.Print(numbers2)
}

func printSlice2(x []int) {
	fmt.Printf("len:%d, cap:%d, slice:%v\n", len(x), cap(x), x)
}
