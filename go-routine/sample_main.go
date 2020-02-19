package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		fmt.Println("inside go routine")
	}()

	msg := <-messages
	fmt.Println(msg)
}
