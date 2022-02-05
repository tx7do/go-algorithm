package sorting

// SelectionSort 选择排序
// @see https://en.wikipedia.org/wiki/Selection_sort
func SelectionSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	i, j, minIndex := 0, 0, 0
	for i = begin; i < end; i++ {
		minIndex = i
		for j = i + 1; j <= end; j++ {
			if array.Less(j, minIndex) {
				minIndex = j
			}
		}
		if minIndex != i {
			array.Swap(i, minIndex)
		}
	}
}
