package sorting

import (
	"math"
)

// IntroSort 自省排序
// @see https://en.wikipedia.org/wiki/Introsort
// @see https://en.wikipedia.org/wiki/Introsort
func IntroSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	depthLimit := 2 * int(math.Round(math.Log2(float64(end))))
	introSort(array, begin, end, depthLimit)
}

func introSort(array Interface, begin, end int, depthLimit int) {
	length := end - begin + 1
	if length < 12 {
		InsertionSort(array, begin, end)
	} else if depthLimit == 0 {
		HeapSort(array, begin, end)
	} else {
		depthLimit -= 1
		pi := quickSortPartition(array, begin, end)

		introSort(array, begin, pi-1, depthLimit)
		introSort(array, pi+1, end, depthLimit)
	}
}
