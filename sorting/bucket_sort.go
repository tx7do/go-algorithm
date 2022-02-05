package sorting

import "math"

const DefaultBucketSize = 5

// BucketSort 桶排序
// @see https://www3.nd.edu/~dchiang/teaching/ds/2015/handouts/bucketsort.pdf
// @see https://www.geeksforgeeks.org/bucket-sort-2/
// @see https://en.wikipedia.org/wiki/Bucket_sort
// @see https://codecrucks.com/bucket-sort/
func BucketSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	minValue, maxValue := array.Get(begin), array.Get(begin)
	i := 0
	for i = begin + 1; i <= end; i++ {
		if Less(array.Get(i), minValue) {
			minValue = array.Get(i)
		}
		if Less(maxValue, array.Get(i)) {
			maxValue = array.Get(i)
		}
	}

	bucketSize := DefaultBucketSize
	bucketCount := int(math.Floor(float64(Minus(maxValue, minValue)/bucketSize))) + 1
	buckets := make([]InterfaceSlice, bucketCount)

	for i = begin; i <= end; i++ {
		bucketIndex := int(math.Floor(float64(Minus(array.Get(i), minValue) / bucketSize)))
		buckets[bucketIndex] = append(buckets[bucketIndex], array.Get(i))
	}

	sortedIndex := 0
	for _, bucket := range buckets {
		InsertionSort(bucket, 0, bucket.Len()-1)
		for _, val := range bucket {
			array.Set(sortedIndex, val)
			sortedIndex++
		}
	}
}
