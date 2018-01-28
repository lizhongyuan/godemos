package main

import "fmt"

func deferRet(x,y int) (z int){
	// defer z += 100  defer must be function call
	defer func() {
		z += 100
	}()
	z = x + y
	return z + 50 // 执行顺序 z = z+50 -> (call defer)z = z+100 -> ret
}

func main() {
	i := deferRet(1,1)
	fmt.Print(i)  // print 152
}
