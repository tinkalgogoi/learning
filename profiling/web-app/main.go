// https://intelligentbee.com/2017/08/01/profiling-web-applications-golang/
package main
import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

// O(n) Fibonacci
func linearFibonacci(n int) int {
	// Create an int array of size n + 1
	v := make([]int, n+1)

	// F(0) = 0
	v[0] = 0
	// F(1) = 1
	v[1] = 1

	// F(i) = F(i-1) + F(i-2)
	for i := 2; i <= n; i++ {
		v[i] = v[i-1] + v[i-2]
	}

	// F(n) - return the n-th Fibonacci number
	return v[n]
}

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

// HTTP request handler
func handler(w http.ResponseWriter, r *http.Request) {
	// return the 25th Fibonacci number in the response payload
	fmt.Fprintf(w, "%d", exponentialFibonacci(45))
}

func main() {
	// Create a new HTTP multiplexer
	mux := http.NewServeMux()

	// Register our handler for the / route
	mux.HandleFunc("/", handler)

	// Add the pprof routes
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	// Start listening on port 8989
	if err := http.ListenAndServe(":8989", mux); err != nil {
		log.Fatal(fmt.Sprintf("Error when starting or running http server: %v", err))
	}
}
