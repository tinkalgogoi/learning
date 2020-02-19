package main

import "fmt"

func main1() {
	s := "cbaeacbbacd"
	p := "abc"
	pMap := make(map[string]bool, len(p))
	for _, v := range p {
		pMap[string(v)] = false
	}
	outerLoop:
	for i, _ := range s {
		for _, v := range p {
			pMap[string(v)] = false
		}
		j:= i
		count := 0
		for j=i; j<i+len(p) && j< len(s); j++ {
			visited, ok := pMap[string(s[j])]
			if ok {
				if visited {
					continue outerLoop
				} else {
					count++
					pMap[string(s[j])] = true
				}
			} else {
				continue
			}
		}
		if count == len(p) {
			fmt.Printf("anagram at index : %d\n", i)
		}
	}
}
func main() {
	//ch := make(chan int)
	//go f(ch)
	//<-ch
	go f2()
	var s string
	fmt.Scan(&s)
	fmt.Println("last main")
}
func f(ch chan int) {
	panic("panicked in f")
	ch <- 0
}
func f2(){
	panic("panicked in f2")
}
