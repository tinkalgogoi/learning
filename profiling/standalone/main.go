package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var wd, _ = os.Getwd()
var cpuprofile = flag.String("cpuprofile", wd+"/profiling/standalone/cpuprofile.out", "write cpu profile to file")

// O(2^n) Fibonacci
func exponentialFibonacci(n int) int {
	// F(0) = 0
	if n == 0 {
		return 0
	}

	// F(1) = 1
	if n == 1 {
		return 1
	}

	// F(n) = F(n-1) + F(n-2) - return the n-th Fibonacci number
	return exponentialFibonacci(n-1) + exponentialFibonacci(n-2)
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(f.Name())
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fib := exponentialFibonacci(45)
	fmt.Println("fib: ",fib)
}
