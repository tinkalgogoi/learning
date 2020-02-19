package main

import "fmt"

func main() {
	count := Solution("BAONXXOLL")
	fmt.Println("count --- ", count)
}

func Solution(S string) int {
	if len(S) < 1 {
		return 0
	}

	m := make(map[rune]int)
	for _, ch := range S {
		if _, exists := m[ch]; !exists {
			m[ch] = 0
		}
		m[ch] = m[ch] +1
	}

	count := 0
	arr := []rune{'A', 'B', 'L', 'O', 'N'}
	arrInt := []int{1, 1, 2, 2, 1}

	constMap := make(map[rune]int)
	for i, _ := range arr {
		if _, exists := constMap[arr[i]]; !exists {
			constMap[arr[i]] = arrInt[i]
		}
	}

	flag := false

	for true {
		for i := 0; i < len(arr); i++ {
			key := arr[i]
			if m[key] >= constMap[key] {
				m[key] = m[key] - constMap[key]
			} else {
				flag = true
				break
			}
		}
		if flag {
			break
		}
		count++
	}
	return count

}


