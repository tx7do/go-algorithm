package sorting

// InterpolationSort 插值排序
// @see https://en.wikipedia.org/wiki/Interpolation_sort
// @see https://cs.colby.edu/courses/F18/cs375/InterpolationSortpaper.pdf
// @see https://wikichi.icu/wiki/Interpolation_sort
// @see http://www.ijcte.org/papers/575-A280.pdf
// @see https://zh.wikipedia.org/wiki/%E6%8F%92%E5%80%BC%E6%8E%92%E5%BA%8F
func InterpolationSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

}
