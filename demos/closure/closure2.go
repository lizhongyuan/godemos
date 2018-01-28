package main

import "fmt"

func main() {
	// 方法1: 每次循环构造一个临时变量 i
	for i := 0; i < 5; i++ {
		i := i
		defer func(){ fmt.Printf("%d ", i) } ()
		// Output: 4 3 2 1 0
	}
	// 方法2: 通过函数参数传参
	for i := 0; i < 5; i++ {
		defer func(i int){ fmt.Printf("%d ", i) } (i)
		// Output: 4 3 2 1 0
	}
}
