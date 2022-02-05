package sorting

import (
	"math"
)

// StoogeSort 臭皮匠排序
// @see https://en.wikipedia.org/wiki/Stooge_sort
func StoogeSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	if array.Less(end, begin) {
		array.Swap(end, begin)
	}

	if length > 2 {
		t := int(math.Floor(float64(end-begin+1) / 3))
		StoogeSort(array, begin, end-t)
		StoogeSort(array, begin+t, end)
		StoogeSort(array, begin, end-t)
	}
}
