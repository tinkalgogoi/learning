package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

func NewNode(d int) *Node {
	return &Node{1,nil,nil}
}

func main(){
	root := NewNode(2)
	root.left = NewNode(1)
	root.right = NewNode(3)

	fmt.Println(root)
}