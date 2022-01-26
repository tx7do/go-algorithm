package sorting

import (
	"math"
)

const DefaultBucketSize = 5

// BucketSort 桶排序
func BucketSort(array IntSlice) {
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

	bucketSize := DefaultBucketSize
	bucketCount := int(math.Floor(float64((maxValue-minValue)/bucketSize))) + 1
	buckets := make([]IntSlice, bucketCount)
	for i = 0; i < length; i++ {
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
