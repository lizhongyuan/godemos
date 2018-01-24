package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Printf("sum: %d\n", sum(1, 2))
	fmt.Printf("mas: %d\n", max(1, 2))
}


