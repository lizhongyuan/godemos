package main

import "fmt"


func Fab(n uint64) (result uint64) {
	if (n > 0) {
		result = n * Fab(n - 1)
		return result
	}
	return 1
}


func main() {
	fmt.Print(Fab(3))
}
