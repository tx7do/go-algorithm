package sorting

import (
	algorithm "go-algorithm"
	"math"
)

// CountingSort 计数排序
func CountingSort(array []int) {
	length := len(array)
	if length < 2 {
		return
	}

	maxValue := array[0]
	for _, val := range array {
		maxValue = algorithm.Max(maxValue, val)
		if val < 0 {
			return
		}
	}

	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)

	for _, val := range array {
		bucket[val]++
	}

	sortedIndex := 0
	for i := 0; i < bucketLen; i++ {
		for bucket[i] > 0 {
			array[sortedIndex] = i
			sortedIndex++
			bucket[i]--
		}
	}
}

// CountingSortNegative 计数排序 - 可以处理负数的版本
func CountingSortNegative(array []int) {
	length := len(array)
	if length < 2 {
		return
	}

	maxValue := array[0]
	minValue := array[0]
	for _, val := range array {
		minValue = algorithm.Min(minValue, val)
		maxValue = algorithm.Max(maxValue, val)
	}

	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)

	negativeBucketLen := 0
	if minValue < 0 {
		negativeBucketLen = int(math.Abs(float64(minValue))) + 1
	}
	negativeBucket := make([]int, negativeBucketLen)

	for _, val := range array {
		if val >= 0 {
			bucket[val]++
		} else {
			negativeBucket[int(math.Abs(float64(val)))]++
		}
	}

	sortedIndex := 0
	for i := negativeBucketLen - 1; i >= 0; i-- {
		for negativeBucket[i] > 0 {
			array[sortedIndex] = -i
			sortedIndex++
			negativeBucket[i]--
		}
	}

	for i := 0; i < bucketLen; i++ {
		for bucket[i] > 0 {
			array[sortedIndex] = i
			sortedIndex++
			bucket[i]--
		}
	}
}
