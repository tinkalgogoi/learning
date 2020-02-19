package main

func main() {

}


func Solution(A []int) int {
	count := 0
	var i,j int
	for i, j = 0, i + 1; j < len(A); {
		for j < len(A) && A[i] == A[j] {
			i++
			j++
		}
		count++
		if i > -1 && j < len(A) {
			if A[i] > A[j] {
				for j < len(A) && A[i] >= A[j] {
					i++
					j++
				}
			} else {
				for j < len(A) && A[i] <= A[j] {
					i++
					j++
				}
			}
			if A[i-1] == A[j-1] {
				count++
			}

		}

	}
	return count
}