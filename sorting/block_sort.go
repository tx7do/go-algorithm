package sorting

// BlockSort 分块排序
// @see https://en.wikipedia.org/wiki/Block_sort
// @see https://www.imsc.res.in/~meena/papers/block-sort.pdf
func BlockSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}
}
