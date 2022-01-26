package sorting

import "math"

// MergeSort 归并排序
// @see https://en.wikipedia.org/wiki/Merge_sort
func MergeSort(array IntSlice) IntSlice {
	length := array.Len()
	if length < 2 {
		return array
	}

	mid := int(math.Floor(float64(length >> 1)))

	left := MergeSort(array[:mid])
	right := MergeSort(array[mid:])

	return merge(left, right)
}

// merge 归并
func merge(left, right IntSlice) IntSlice {
	lengthLeft := left.Len()
	lengthRight := right.Len()

	var result = make(IntSlice, 0, lengthLeft+lengthRight)

	i, j := 0, 0
	for i < lengthLeft && j < lengthRight {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
