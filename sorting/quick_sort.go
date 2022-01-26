package sorting

import "math/rand"

// QuickSort 快速排序
func QuickSort(array []int) {
	length := len(array)
	if length < 2 {
		return
	}

	low, high := 0, length-1

	pivot := rand.Int() % len(array)

	array[pivot], array[high] = array[high], array[pivot]

	for i := range array {
		if array[i] < array[high] {
			array[low], array[i] = array[i], array[low]
			low++
		}
	}

	array[low], array[high] = array[high], array[low]

	QuickSort(array[:low])
	QuickSort(array[low+1:])
}
