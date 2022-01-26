package sorting

import (
	"math"
)

// InPlaceMergeSort 原地归并排序
// @see https://en.wikipedia.org/wiki/Merge_sort
func InPlaceMergeSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	mid := int(math.Floor(float64(length >> 1)))

	InPlaceMergeSort(array[:mid])
	InPlaceMergeSort(array[mid:])

	inPlaceMerge(array, mid)
}

// inPlaceMerge 原地归并
func inPlaceMerge(array IntSlice, mid int) {
	indexLeft, indexRight := 0, mid
	endLeft, endRight := mid-1, array.Len()-1

	if array[endLeft] <= array[indexRight] {
		return
	}

	sortedIndex, tempValue := 0, 0
	for indexLeft <= endLeft && indexRight <= endRight {
		if array[indexLeft] <= array[indexRight] {
			indexLeft++
		} else {
			tempValue = array[indexRight]
			sortedIndex = indexRight

			for sortedIndex != indexLeft {
				array[sortedIndex] = array[sortedIndex-1]
				sortedIndex--
			}
			array[indexLeft] = tempValue

			indexLeft++
			endLeft++
			indexRight++
		}
	}
}
