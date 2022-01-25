package sort

import "math"

// MergeSort 归并排序
func MergeSort(array []int) []int {
	length := len(array)
	if length < 2 {
		return array
	}

	mid := int(math.Floor(float64(length >> 1)))
	first := MergeSort(array[:mid])
	second := MergeSort(array[mid:])
	return merge(first, second)
}

func merge(left, right []int) []int {
	var result = make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
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
