package sorting

import "math"

// ProxmapSort 相邻图排序
// @see https://en.wikipedia.org/wiki/Proxmap_sort
// @see https://www.cs.uah.edu/~rcoleman/CS221/Sorting/ProxMapSort.html
// @see https://www.ics.uci.edu/~jacobson/ics23/ProxmapHandout.pdf
func ProxmapSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	i := 0

	minValue, maxValue := array[begin], array[begin]
	for i = begin + 1; i <= end; i++ {
		if array[i] < minValue {
			minValue = array[i]
		}
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}

	flagArray := make([]bool, length)
	sortedArray := make(IntSlice, length)
	interpolation := make(IntSlice, length)
	hitCount := make(IntSlice, length)

	for i = begin; i <= end; i++ {
		interpolation[i] = int(math.Floor(float64(((array[i] - minValue) / (maxValue - minValue)) * (end - begin))))
		hitCount[interpolation[i]]++
	}
	hitCount[end] = length - hitCount[end]
	for i = end; i > begin; i-- {
		hitCount[i-1] = hitCount[i] - hitCount[i-1]
	}

	var insertIndex = 0
	var insertStart = 0
	for i = begin; i <= end; i++ {
		insertIndex = hitCount[interpolation[i]]
		insertStart = insertIndex
		for flagArray[insertIndex] {
			insertIndex++
		}
		for insertIndex > insertStart && array[i] < sortedArray[insertIndex-1] {
			sortedArray.Swap(insertIndex, insertIndex-1)
			flagArray[insertIndex-1] = true
			flagArray[insertIndex] = true
			insertIndex--
		}
		sortedArray[insertIndex] = array[i]
		flagArray[insertIndex] = true
	}
	for i = begin; i <= end; i++ {
		array[i] = sortedArray[i]
	}
}
