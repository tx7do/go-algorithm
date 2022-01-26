package sorting

// SelectionSort 选择排序
func SelectionSort(array []int) {
	length := len(array)
	if length < 2 {
		return
	}

	i, j := 0, 0
	minIndex := 0
	for i = 0; i < length-1; i++ {
		minIndex = i
		for j = i + 1; j < length; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}
