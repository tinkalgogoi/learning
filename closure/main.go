package main

import (
"fmt"
)

func makeEvenGenerator() func() int {
	i := 2
	return func() (ret int) {
		ret = i
		i += 2
		return ret
	}
}

func main() {
	nextEven := makeEvenGenerator() // makeEvenGenerator runs once here
	fmt.Println(nextEven()) // nextEven is not makeEvenGenerator and it is a type (func() int), so each time return section in
	// makeEvenGenerator gets executed
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

