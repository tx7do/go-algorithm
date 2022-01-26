package sorting

import (
	"math"
)

// CountingSort 计数排序
func CountingSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	maxValue := array[0]
	i := 0
	for i = 1; i < length; i++ {
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}

	bucketLen := maxValue + 1
	bucket := make(IntSlice, bucketLen)

	for _, val := range array {
		bucket[val]++
	}

	sortedIndex := 0
	for i = 0; i < bucketLen; i++ {
		for bucket[i] > 0 {
			array[sortedIndex] = i
			sortedIndex++
			bucket[i]--
		}
	}
}

// CountingSortNegative 计数排序 - 可以处理负数的版本
func CountingSortNegative(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	minValue, maxValue := array[0], array[0]
	i := 0
	for i = 1; i < length; i++ {
		if array[i] < minValue {
			minValue = array[i]
		}
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}

	bucketLen := maxValue + 1
	bucket := make(IntSlice, bucketLen)

	negativeBucketLen := 0
	if minValue < 0 {
		negativeBucketLen = int(math.Abs(float64(minValue))) + 1
	}
	negativeBucket := make(IntSlice, negativeBucketLen)

	for _, val := range array {
		if val >= 0 {
			bucket[val]++
		} else {
			negativeBucket[int(math.Abs(float64(val)))]++
		}
	}

	sortedIndex := 0
	for i = negativeBucketLen - 1; i >= 0; i-- {
		for negativeBucket[i] > 0 {
			array[sortedIndex] = -i
			sortedIndex++
			negativeBucket[i]--
		}
	}

	for i = 0; i < bucketLen; i++ {
		for bucket[i] > 0 {
			array[sortedIndex] = i
			sortedIndex++
			bucket[i]--
		}
	}
}
