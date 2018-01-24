package main

import "fmt"

func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("a变量的地址是:%x\n", ip)

	fmt.Printf("变量指针访问:%d\n", *ip)
}
