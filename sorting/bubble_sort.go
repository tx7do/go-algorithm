package sorting

// BubbleSort 冒泡排序
func BubbleSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	i, j := 0, 0
	for i = begin; i < end; i++ {
		for j = begin; j < end-i; j++ {
			if array.Less(j+1, j) {
				array.Swap(j, j+1)
			}
		}
	}
}
