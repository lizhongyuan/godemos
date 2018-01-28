package main

import "fmt"

var ch chan int = make(chan int, 10)	// 长度为10的channel
//var ch chan int = make(chan int)		// 如果不立刻消费就会死锁

func foo(id int) { //id: 这个routine的标号
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)	// 5个goroutine, 顺序打乱
		//ch <- i			//
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Print(<- ch)
	}
}
