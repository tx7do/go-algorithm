package sorting

// GnomeSort 侏儒排序
// @see https://zh.wikipedia.org/wiki/%E4%BE%8F%E5%84%92%E6%8E%92%E5%BA%8F
func GnomeSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	i := 0
	for i < length {
		if i == 0 {
			i++
		}
		if array[i] >= array[i-1] {
			i++
		} else {
			//array[i], array[i-1] = array[i-1], array[i]
			array.Swap(i, i-1)
			i--
		}
	}
}
