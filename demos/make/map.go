package main

import "fmt"

func main() {
	var curMap map[int]int = make(map[int]int)

	curMap[0] = 0
	curMap[1] = 1

	var curArr []int = make([]int, 5)

	fmt.Print(curMap)
	fmt.Print(curArr)
}


