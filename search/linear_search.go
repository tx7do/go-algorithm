package search

// LinearSearch 顺序查找
func LinearSearch(array IntSlice, target int) int {
	length := array.Len()
	if length == 0 {
		return -1
	}

	for i := 0; i < length; i++ {
		if array[i] == target {
			return i
		}
	}

	return -1
}
