package search

// LinearSearch 顺序查找
func LinearSearch(array []int, target int) int {
	if len(array) == 0 {
		return -1
	}
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			return i
		}
	}
	return -1
}
