package main

import "fmt"

type Node struct
{
	data int
	next *Node
}

// LinkedList is not a way in geekforgeeks but we require it to use a single
// reference to a particular linkedLIst object with methods
//
// Or instead of this we can use normal method as geeksforgeeks do for C
// where each method we have to pass head reference
type LinkedList struct {
	head *Node
}
// example a method in LinkedList type
func (ll *LinkedList) AddToHead(data int) {
	ll.head = &Node{data, ll.head}
}

func main() {
	var head *Node
	//Create new node
	head = &Node{1, nil}
	fmt.Println(head)

	// using "for" as "while" as in geekforgeeks
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	//var prev *Node
	//prev = head
	//fmt.Println(prev)



}
