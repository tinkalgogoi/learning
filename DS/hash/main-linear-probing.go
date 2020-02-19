// https://www.sanfoundry.com/c-program-implement-hash-tables-linear-probing/
package main

import "fmt"

const sizee = 7

var HashTable [sizee]int

func init() {
	for i := range HashTable {
		HashTable[i] = -1
	}
}
func hashCodee(val int) (key int) {
	//calculate hash key
	key = val % sizee
	return key
}

func insertt(val int) {

	key := hashCodee(val)

	for HashTable[key] != -1 {
		if key == len(HashTable)-1 {
			fmt.Println("no space left")
			return
		}
		if HashTable[key] == val {
			fmt.Println("val already present : ", val)
			return
		}

		key = hashCodee(key + 1) // linear probing
	}

	HashTable[key] = val

}

//TODO:  add search method.
// t- how to get the exact value when collision happens i.e different key generate same index.

func main() {

	insertt(7)
	insertt(0)
	insertt(3)
	insertt(10)
	insertt(4)
	insertt(5)
	insertt(1)
	insertt(2)
	insertt(2089)

	for _, v := range HashTable {
		fmt.Printf(" %d", v)
	}
}
