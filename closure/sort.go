package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 11, -5, 8, 2, 0, 12}
	sort.Ints(numbers)
	fmt.Println("Sorted:", numbers)

	index := sort.Search(len(numbers), mySort(numbers))
	fmt.Println("The first number >= 7 is at index:", index)
	fmt.Println("The first number >= 7 is:", numbers[index])
}

func mySort(numbers []int) func(i int) bool {
	return func(i int) bool {
		return numbers[i] >= 7  //This closure has access to the numbers slice
	}
}