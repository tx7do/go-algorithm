package sorting

import "math/rand"

// QuickSort 快速排序
func QuickSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	low, high := 0, length-1

	pivot := rand.Int() % array.Len()
	array.Swap(pivot, high)
	//array[pivot], array[high] = array[high], array[pivot]

	for i := range array {
		if array[i] < array[high] {
			array.Swap(low, i)
			//array[low], array[i] = array[i], array[low]
			low++
		}
	}

	//array[low], array[high] = array[high], array[low]
	array.Swap(low, high)

	QuickSort(array[:low])
	QuickSort(array[low+1:])
}
