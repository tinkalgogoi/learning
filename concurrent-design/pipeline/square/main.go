package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				fmt.Println("returning from gen")
				return
			}
		}
		close(out)
	}()
	return out
}

// t- sq is multiple function (or workers) reading from same channel (fan out here)
func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				fmt.Println("returning from sq")
				return
			}
		}
	}()
	return out
}

// t- merge fan in sq workers individual channel to a single channel
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c or done is closed, then calls
	// wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				fmt.Println("returning from 1 of the output routine in merge")
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})

	in := gen(done, 2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(done, in)
	c2 := sq(done, in)

	// Consume the first value from output.
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9
	close(done)



	// exit main after some task
	exit := make(chan struct{})
	go func() {
		sec := time.Duration(10)
		fmt.Println("sleeping for sec : ", sec)
		time.Sleep(sec * time.Second)
		exit <- struct{}{}
	}()
	<-exit
	fmt.Println("exiting from main")
}
