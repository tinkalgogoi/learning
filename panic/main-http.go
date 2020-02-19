package main

import (
"fmt"
"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}
type H struct{
	name string
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	//panic("boom")

	fmt.Println("hello b4 panicking")
	// panic on nil pointer
	var h *H
	h = nil
	fmt.Print(h.name)

}
