package sorting

// InsertionSort 插入排序
func InsertionSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	i, j := begin+1, 0
	for ; i <= end; i++ {
		for j = i; j > 0 && array.Less(j, j-1); j-- {
			array.Swap(j, j-1)
		}
	}
}
