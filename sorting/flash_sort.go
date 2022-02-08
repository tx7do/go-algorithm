package sorting

import "math"

// FlashSort 闪电排序
// @see https://en.wikipedia.org/wiki/Flashsort
func FlashSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	minValue, maxValue := array[begin], array[begin]
	i := 0
	for i = begin + 1; i <= end; i++ {
		if array[i] < minValue {
			minValue = array[i]
		}
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}

	if minValue == maxValue {
		return
	}

	var start = 0
	var high = length
	var size = end - start

	var hitCount = make(IntSlice, high)

	var interpolation = 0
	for i = start; i < high; i++ {
		interpolation = int(math.Floor(float64(((array[i] - minValue) / (maxValue - minValue)) * (size - 1))))
		hitCount[interpolation]++
	}
	hitCount[0] = hitCount[0] - 1
	for i = 1; i < high; i++ {
		hitCount[i] = hitCount[i] + hitCount[i-1]
	}

	var tag = make([]bool, high)
	for i = start; i < high; i++ {
		for tag[i] == false {
			interpolation = int(math.Floor(float64(((array[i] - minValue) / (maxValue - minValue)) * (size - 1))))
			tag[hitCount[interpolation]] = true
			array.Swap(i, hitCount[interpolation])
			hitCount[interpolation]--
		}
	}

	InsertionSort(array, begin, end)
}
