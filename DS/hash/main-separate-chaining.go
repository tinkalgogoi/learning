// modified from https://www.log2base2.com/algorithms/searching/open-hashing.html
// same as https://www.geeksforgeeks.org/hashing-set-2-separate-chaining/
package main

import "fmt"

type node struct {
	data int
	next *node
}

const size = 7
var Chain [size]*node

func hashCode(val int) (key int){
	//calculate hash key
	key = val % size
	return key
}

func insert(value int) {
	//create a new node with value
	newNode := &node{value, nil}

	//get hash key
	key := hashCode(value)

	//check if chain[key] is empty
	if Chain[key] == nil {
		Chain[key] = newNode
	} else { //collision
		//add the node at the end of chain[key].
		temp := Chain[key]
		for temp.next != nil {
			temp = temp.next
		}

		temp.next = newNode
	}
}

//TODO:  add search method.
// how to get the exact value when collision happens i.e different key generate same index.
// t- may be while inserting wrap the value with corresponding key also(say a struct to hold both). Then traverse linkedlist to see the key and get value.
// so we insert keys and values both.
// https://www.geeksforgeeks.org/implementing-our-own-hash-table-with-separate-chaining-in-java/



func print() {

	for i := 0; i < size; i++ {
		temp := Chain[i]
		fmt.Printf("chain[%d] ----> ", i)
		for temp != nil {
			fmt.Printf("%d -->", temp.data)
			temp = temp.next
		}
		fmt.Println("nil")
	}
}

func main() {

	insert(7)
	insert(0)
	insert(3)
	insert(10)
	insert(4)
	insert(5)
	insert(1)
	insert(2)
	insert(2089)

	print()

}
