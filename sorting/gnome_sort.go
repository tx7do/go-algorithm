package sorting

// GnomeSort 侏儒排序
// @see https://zh.wikipedia.org/wiki/%E4%BE%8F%E5%84%92%E6%8E%92%E5%BA%8F
func GnomeSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	for i := begin; i <= end; {
		if i == 0 {
			i++
		}
		if array.Less(i, i-1) {
			array.Swap(i, i-1)
			i--
		} else {
			i++
		}
	}
}
