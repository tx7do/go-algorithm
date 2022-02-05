package sorting

import (
	"math"
)

// SlowSort 慢速排序
// @see https://www2.kenyon.edu/Depts/Math/Aydin/Teach/Sp15/218/SlowSorting.pdf
// @see https://en.wikipedia.org/wiki/Slowsort
func SlowSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	mid := int(math.Floor(float64(begin + (end-begin)>>1)))

	if length > 2 {
		SlowSort(array, begin, mid)
		SlowSort(array, mid+1, end)
	}

	if array.Less(end, mid) {
		array.Swap(end, mid)
	}

	if end > 1 {
		SlowSort(array, begin, end-1)
	}
}
