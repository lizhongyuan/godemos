package main

import "fmt"

var channel chan int = make(chan int)

func emit0() {
	channel <- 0
}

func main() {
	go emit0()
	fmt.Println(<-channel)
}
