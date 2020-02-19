// modified and taken from
// https://code.tutsplus.com/tutorials/context-based-programming-in-go--cms-29290
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func contextCancelDemo(ctx context.Context) error {
	deadline, ok := ctx.Deadline()
	name := ctx.Value("name")
	if ok {
		fmt.Println(name, "will expire at:", deadline)
	} else {
		fmt.Println(name, "has no deadline")
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "Context has been cancelled from main...")
			return ctx.Err()
		default:
		}
	}
}

func contextTimeoutDemo(ctx context.Context) error {
	deadline, ok := ctx.Deadline()
	name := ctx.Value("name")
	if ok {
		fmt.Println(name, "will expire at:", deadline)
	} else {
		fmt.Println(name, "has no deadline")
	}

	for {
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			fmt.Println(name, "Context has been cancelled as timeout reached...")
			return ctx.Err()
		default:
		}
	}
}

func contextDeadlineDemo(ctx context.Context) error {
	deadline, ok := ctx.Deadline()
	name := ctx.Value("name")
	if ok {
		fmt.Println(name, "will expire at:", deadline)
	} else {
		fmt.Println(name, "has no deadline")
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "Context has been cancelled as deadline reached...")
			return ctx.Err()
		default:
		}
	}
}

func main() {
	timeout := 5 * time.Second
	deadline := time.Now().Add(10 * time.Second)
	timeOutContext, _ := context.WithTimeout(
		context.Background(), timeout)
	cancelContext, cancelFunc := context.WithCancel(
		context.Background())
	// t- Q- How WithDeadline automatically closes done channel on deadline expires? Should have
	// been ticker running in some go routine but couldn't find it.
	// Ans- time.AfterFunc calls d argument func u provide when timer expires but still couldn't find how its updating the ticker.
	// example of time.AfterFunc - https://play.golang.org/p/DoYuddZYSXR
	deadlineContext, _ := context.WithDeadline(
		context.Background(), deadline)

	go contextCancelDemo(context.WithValue(
		cancelContext, "name", "[cancelContext]"))

	go contextTimeoutDemo(context.WithValue(
		timeOutContext, "name", "[timeoutContext]"))

	go contextDeadlineDemo(context.WithValue(
		deadlineContext, "name", "[deadlineContext]"))

	// This will cancel the deadline context as well as its
	// child - the cancelContext
	fmt.Println("Cancelling the cancel context...")
	cancelFunc()

	// enter to exit from main
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
