package main

import "fmt"

func adder() (func(int) int) {
	sum := 0
	return func(a int) int {
		sum += a
		return sum
	}
}

func main() {
	var add2 = adder()
	var add10 = adder()

	for i := 0; i < 10; i++ {
		fmt.Println(add2(2))
		fmt.Println(add10(10))
	}
}
