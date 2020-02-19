// taken from
// https://code.tutsplus.com/tutorials/context-based-programming-in-go--cms-29290

package main

import (
	"context"
	"fmt"
	"time"
)

func contextDemo(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	name := ctx.Value("name")
	for {
		if ok {
			fmt.Println(name, "will expire at:", deadline)
		} else {
			fmt.Println(name, "has no deadline")
		}
		time.Sleep(time.Second)
	}
}

func main() {
	timeout := 3 * time.Second
	deadline := time.Now().Add(4 * time.Hour)
	timeOutContext, _ := context.WithTimeout(
		context.Background(), timeout)
	cancelContext, cancelFunc := context.WithCancel(
		context.Background())
	deadlineContext, _ := context.WithDeadline(
	    // Note- adding cancelContext here as parent
		cancelContext, deadline)

	go contextDemo(context.WithValue(
		timeOutContext, "name", "[timeoutContext]"))
	go contextDemo(context.WithValue(
		cancelContext, "name", "[cancelContext]"))
	go contextDemo(context.WithValue(
		deadlineContext, "name", "[deadlineContext]"))

	// Wait for the timeout to expire
	<-timeOutContext.Done()

	// This will cancel the deadline context as well as its
	// child - the cancelContext
	fmt.Println("Cancelling the cancel context...")
	cancelFunc()

	<-cancelContext.Done()
	fmt.Println("The cancel context has been cancelled...")

	// Wait for both contexts to be cancelled
	<-deadlineContext.Done()  // t- |z this is supposed to get call when deadline time comes
	fmt.Println("The deadline context has been cancelled...")
}
