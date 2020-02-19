// https://golang.org/doc/effective_go.html#allocation_new
package main

import "fmt"

type S1 struct {
	i int
}
func main()  {
	slice1 := make([]int, 0)
	fmt.Println("len n cap slice1", len(slice1), cap(slice1))
	s1 := new(S1)
	fmt.Println(s1)
	m1 := make(map[string]int)
	fmt.Println("len m1 :",len(m1))
	m1["tin"] = 1
	fmt.Println(m1)
	m := new(map[string]string)
	m2 := *m
	m2["tin"] = "go"

	fmt.Println(m2)
}


