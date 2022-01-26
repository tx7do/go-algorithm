package sorting

import (
	algorithm "go-algorithm"
	"math"
)

const DefaultBucketSize = 5

// BucketSort 桶排序
func BucketSort(array []int) {
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

	bucketSize := DefaultBucketSize
	bucketCount := int(math.Floor(float64((maxValue-minValue)/bucketSize))) + 1
	buckets := make([][]int, bucketCount)
	for i := 0; i < length; i++ {
		bucketIndex := int(math.Floor(float64((array[i] - minValue) / bucketSize)))
		buckets[bucketIndex] = append(buckets[bucketIndex], array[i])
	}

	sortedIndex := 0
	for _, bucket := range buckets {
		InsertionSort(bucket)
		for _, val := range bucket {
			array[sortedIndex] = val
			sortedIndex++
		}
	}
}
