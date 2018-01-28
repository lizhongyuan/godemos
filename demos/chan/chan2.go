package main

import "fmt"

func main() {
	channel := make(chan string)

	go func(message string) {
		channel <- message
	}("Ping")

	fmt.Println(<-channel)
}
