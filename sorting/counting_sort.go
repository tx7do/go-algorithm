package sorting

import (
	"math"
)

// CountingSort 计数排序
func CountingSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	maxValue := array[begin]
	i := 0
	for i = begin + 1; i <= end; i++ {
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}

	bucketLen := maxValue + 1
	bucket := make(IntSlice, bucketLen)

	for i = begin; i <= end; i++ {
		bucket[array[i]]++
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
func CountingSortNegative(array IntSlice, begin, end int) {
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

	bucketLen := maxValue + 1
	bucket := make(IntSlice, bucketLen)

	negativeBucketLen := 0
	if minValue < 0 {
		negativeBucketLen = int(math.Abs(float64(minValue))) + 1
	}
	negativeBucket := make(IntSlice, negativeBucketLen)

	for i = begin; i <= end; i++ {
		if array[i] >= 0 {
			bucket[array[i]]++
		} else {
			negativeBucket[int(math.Abs(float64(array[i])))]++
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
