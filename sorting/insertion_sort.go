package sorting

// InsertionSort 插入排序
func InsertionSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	i, preIndex, current := 0, 0, 0
	for i = 1; i < length; i++ {
		preIndex = i - 1
		current = array[i]
		for preIndex >= 0 && array[preIndex] > current {
			array[preIndex+1] = array[preIndex]
			preIndex--
		}
		array[preIndex+1] = current
	}
}
