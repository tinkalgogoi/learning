//This is a LRU cache
//https://github.com/Lebonesco/go_lru_cache/blob/master/main.go

// Reading - https://medium.com/@krishankantsinghal/my-first-blog-on-medium-583159139237
package main

import (
	"fmt"
)

const SIZE = 5 // size of cache

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

// double linked list
type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

// maps string to node in Queue
type Hash map[string]*Node

type Cache struct {
	queue Queue
	hash  Hash
}

func NewCache() Cache {
	return Cache{queue: NewQueue(), hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}
//t- Check is the method who has the main logic of LRU
func (c *Cache) Check(str string) {
	node := &Node{}
	//t- if node is available in cache remove it
	if foundNode, ok := c.hash[str]; ok {
		node = c.Remove(foundNode)
	} else {
		//t- else its new node (ex- say fetch it from database)
		node = &Node{Val: str}
	}
	//t- add the removed/new node in front of queue
	c.AddToHead(node)
	//t- add to hash
	c.hash[str] = node
}

//t- removes node n at any location
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)
	left := n.Left
	right := n.Right
	left.Right = right
	right.Left = left
	c.queue.Length -= 1

	delete(c.hash, n.Val)
	return n
}

//t- add at the head and while adding if length increases than SIZE than remove the tail(least used)
func (c *Cache) AddToHead(n *Node) {
	fmt.Printf("add: %s\n", n.Val)
	tmp := c.queue.Head.Right
	c.queue.Head.Right = n
	n.Left = c.queue.Head
	n.Right = tmp
	tmp.Left = n

	//t- removing the least recently used when queue is full
	c.queue.Length++
	if c.queue.Length > SIZE {
		c.Remove(c.queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf(" <--> ")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"cat", "blue", "dog", "tree", "dragon",
		"potato", "house", "tree", "cat"} {
		cache.Check(word)
		cache.Display()
	}
}
