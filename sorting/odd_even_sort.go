package sorting

// OddEvenSort 奇偶排序
// @see https://en.wikipedia.org/wiki/Odd–even_sort
func OddEvenSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	sorted := false
	i := 0
	for !sorted {
		sorted = true
		for i = begin + 1; i < end; i += 2 {
			if array.Less(i+1, i) {
				array.Swap(i, i+1)
				sorted = false
			}
		}
		for i = begin; i < end; i += 2 {
			if array.Less(i+1, i) {
				array.Swap(i, i+1)
				sorted = false
			}
		}
	}
}
