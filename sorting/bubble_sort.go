package sorting

// BubbleSort 冒泡排序
func BubbleSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	i, j := 0, 0
	for i = 0; i < length-1; i++ {
		for j = 0; j < length-1-i; j++ {
			if array[j] > array[j+1] {
				//array[j+1], array[j] = array[j], array[j+1]
				array.Swap(j, j+1)
			}
		}
	}
}