package sorting

import "math"

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

// InPlaceMergeSort 原地归并排序
// @see https://en.wikipedia.org/wiki/Merge_sort
// @see https://www.geeksforgeeks.org/in-place-merge-sort/
func InPlaceMergeSort(array Interface, begin, end int) {
	if begin < end {
		mid := int(math.Floor(float64(begin + (end-begin)>>1)))

		InPlaceMergeSort(array, begin, mid)
		InPlaceMergeSort(array, mid+1, end)

		inPlaceMerge(array, begin, mid, end)
	}
}

// inPlaceMerge 原地归并
func inPlaceMerge(array Interface, begin, mid, end int) {
	indexLeft, indexRight := begin, mid+1
	endLeft, endRight := mid, end

	if array.Less(endLeft, indexRight) {
		return
	}

	sortedIndex := 0
	for indexLeft <= endLeft && indexRight <= endRight {
		if array.Less(indexLeft, indexRight) {
			indexLeft++
		} else {
			tempValue := array.Get(indexRight)
			sortedIndex = indexRight

			for sortedIndex != indexLeft {
				array.Swap(sortedIndex, sortedIndex-1)
				sortedIndex--
			}
			array.Set(indexLeft, tempValue)

			indexLeft++
			endLeft++
			indexRight++
		}
	}
}
