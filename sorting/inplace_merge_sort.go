package sorting

import (
	"math"
)

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
