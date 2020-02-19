package main

func partition(arr[] int, left int, right int) int {
	i := left
	j := right - 1
	pivot := arr[right]

	for i<=j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}

		if i<= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	arr[i], pivot = pivot, arr[i]

	return i
}


