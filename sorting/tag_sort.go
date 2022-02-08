package sorting

// TagSort 标签排序
// @see https://www.geeksforgeeks.org/tag-sort/
func TagSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	tag := make(IntSlice, length)
	for i := 0; i < length; i++ {
		tag[i] = i
	}

	i, j := 0, 0
	for i = 0; i < length; i++ {
		for j = i + 1; j < length; j++ {
			if array.Less(tag[j], tag[i]) {
				array.Swap(i, j)
			}
		}
	}
}
