package sorting

import (
	"fmt"
	"math"
)

// InterpolationSort 插值排序
// @see https://en.wikipedia.org/wiki/Interpolation_sort
// @see https://cs.colby.edu/courses/F18/cs375/InterpolationSortpaper.pdf
// @see https://wikichi.icu/wiki/Interpolation_sort
// @see http://www.ijcte.org/papers/575-A280.pdf
// @see https://zh.wikipedia.org/wiki/%E6%8F%92%E5%80%BC%E6%8E%92%E5%BA%8F
func InterpolationSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	i := 0

	high := length
	bucketSize := make(IntSlice, 0)
	bucketSize = append(bucketSize, high)
	for bucketSize.Len() > 0 {
		size := bucketSize[bucketSize.Len()-1]
		start := high - size

		bucketSize = bucketSize[:bucketSize.Len()-1]

		minValue, maxValue := array[start], array[start]
		for i = start + 1; i < high; i++ {
			if array[i] < minValue {
				minValue = array[i]
			}
			if array[i] > maxValue {
				maxValue = array[i]
			}
		}

		if minValue == maxValue {
			high = high - size
		} else {
			var interpolation = 0
			var bucket = make([]IntSlice, size)
			for i = 0; i < size; i++ {
				interpolation = int(math.Floor(float64(((array[i] - minValue) / (maxValue - minValue)) * (size - 1))))
				bucket[interpolation] = append(bucket[interpolation], array[i])
			}
			for i = 0; i < size; i++ {
				if bucket[i].Len() > 0 {
					for j := 0; j < bucket[i].Len(); j++ {
						array[start] = bucket[i][j]
						start++
					}
					bucketSize = append(bucketSize, bucket[i].Len())
				}
			}
		}
	}
}

// InPlaceInterpolationSort 原地插值排序
func InPlaceInterpolationSort(array IntSlice, begin, end int) {
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

	if maxValue == minValue {
		return
	}

	position := length
	for i = 0; i < length; i++ {
		position = int(math.Floor(float64(((array[i] - minValue) / (maxValue - minValue)) * (length - 1))))
		for i != position {
			array.Swap(i, position)
			fmt.Println(i, position, array[i], array[position])
			position = int(math.Floor(float64(((array[i] - minValue) / (maxValue - minValue)) * (length - 1))))
		}
	}
}

// InterpolationTagSort 插值标签排序
func InterpolationTagSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	i := 0

	var start = 1
	var high = length
	var Tag = make([]bool, high)
	Tag[0] = true

	for high > 0 {
		for start > 0 {
			start--
			if Tag[start] == false {
				break
			}
		}

		minValue, maxValue := array[start], array[start]
		for i = start + 1; i < high; i++ {
			if array[i] < minValue {
				minValue = array[i]
			}
			if array[i] > maxValue {
				maxValue = array[i]
			}
		}
		if minValue == maxValue {
			high = start
		} else {
			var interpolation = 0
			var size = high - start
			var bucket = make([]IntSlice, size)
			for i = 0; i < size; i++ {
				interpolation = int(math.Floor(float64(((array[i] - minValue) / (maxValue - minValue)) * (size - 1))))
				bucket[interpolation] = append(bucket[interpolation], array[i])
			}
			for i = 0; i < size; i++ {
				if bucket[i].Len() > 0 {
					Tag[start] = true
					for j := 0; j < bucket[i].Len(); j++ {
						array[start] = bucket[i][j]
						start++
					}
				}
			}
		}
	}
}

// InterpolationInsertSort 插值排序混合插入排序
func InterpolationInsertSort(array IntSlice, begin, end int) {
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

	if maxValue == minValue {
		return
	}

	var start = 0
	var size = length
	var bucket = make([]IntSlice, size)
	var buckets = size - 1
	var rgn = maxValue - minValue
	var interpolation = 0
	var insertIndex = 0
	var k = 0
	for i = 0; i < size; i++ {
		interpolation = int(math.Floor(float64(((array[i] - minValue) / rgn) * buckets)))
		bucket[interpolation] = append(bucket[interpolation], array[i])
		insertIndex = 0
		for array[i] > bucket[interpolation][insertIndex] {
			insertIndex++
		}
		k = bucket[interpolation].Len() - 1
		for k > insertIndex {
			bucket[interpolation][k] = bucket[interpolation][k-1]
			k--
		}
		bucket[interpolation][insertIndex] = array[i]
	}
	for i = 0; i < size; i++ {
		for j := 0; j < bucket[i].Len(); j++ {
			array[start] = bucket[i][j]
			start++
		}
	}
}

// HistogramSort 直方图排序
func HistogramSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	var high = length
	var sortedArray = make(IntSlice, high)
	var interpolation = make(IntSlice, high)
	var hitCount = make(IntSlice, high)

	i := 0
	var divideSize = make(IntSlice, 0)
	divideSize = append(divideSize, high)
	for divideSize.Len() > 0 {
		size := divideSize[divideSize.Len()-1]
		divideSize = divideSize[:divideSize.Len()-1]

		start := high - size
		minValue, maxValue := array[start], array[start]
		for i = start + 1; i < high; i++ {
			if array[i] < minValue {
				minValue = array[i]
			}
			if array[i] > maxValue {
				maxValue = array[i]
			}
		}

		if minValue == maxValue {
			high = high - size
		} else {
			for i = start; i < high; i++ {
				interpolation[i] = start + int(math.Floor(float64(((array[i]-minValue)/(maxValue-minValue))*(size-1))))
				hitCount[interpolation[i]]++
			}
			for i = start; i < high; i++ {
				if hitCount[i] > 0 {
					divideSize = append(divideSize, hitCount[i])
				}
			}
			hitCount[high-1] = high - hitCount[high-1]
			for i = high - 1; i > start; i-- {
				hitCount[i-1] = hitCount[i] - hitCount[i-1]
			}
			for i = start; i < high; i++ {
				sortedArray[hitCount[interpolation[i]]] = array[i]
				hitCount[interpolation[i]]++
			}
			for i = start; i < high; i++ {
				array[i] = sortedArray[i]
			}
		}
	}

}
