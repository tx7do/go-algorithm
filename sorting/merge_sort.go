package sorting

// MergeSort 归并排序
// @see https://en.wikipedia.org/wiki/Merge_sort
// @see https://www.enjoyalgorithms.com/blog/merge-sort-algorithm
// @see https://qvault.io/golang/merge-sort-golang/
func MergeSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	tmp := mergeSort(array.Part(begin, end+1))
	for i := 0; i < length; i++ {
		array.Set(i, tmp.Get(i))
	}
}

func mergeSort(array Interface) Interface {
	if array.Len() < 2 {
		return array
	}

	mid := divisionByTwo(array.Len())

	left := mergeSort(array.Part(0, mid))
	right := mergeSort(array.Part(mid, array.Len()))

	return spaceMerge(left, right)
}

// spaceMerge 归并
func spaceMerge(left, right Interface) Interface {
	lengthLeft := left.Len()
	lengthRight := right.Len()

	var result = make(InterfaceSlice, 0, lengthLeft+lengthRight)

	i, j := 0, 0
	for i < lengthLeft && j < lengthRight {
		if less(left.Get(i), right.Get(j)) {
			result = append(result, left.Get(i))
			i++
		} else {
			result = append(result, right.Get(j))
			j++
		}
	}

	for ; i < lengthLeft; i++ {
		result = append(result, left.Get(i))
	}
	for ; j < lengthRight; j++ {
		result = append(result, right.Get(j))
	}

	return result
}
