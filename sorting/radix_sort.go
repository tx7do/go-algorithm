package sorting

import (
	"fmt"
	"math"
)

type Buckets []IntSlice

// RadixSort 基数排序
// @see https://mp.weixin.qq.com/s/Z8gU9QLpMnA-zoMc9ZeR2w
// @see https://www.geeksforgeeks.org/radix-sort/
// @see https://en.wikipedia.org/wiki/Radix_sort
func RadixSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	maxNumber := getMaxNumber(array, begin, end)

	const NumberOfBuckets = 10

	n := 1
	bucket := make(Buckets, NumberOfBuckets)
	for n <= maxNumber {
		for _, v := range array {
			digit := getDigit(v, n)
			bucket[digit] = append(bucket[digit], v)
		}
		n *= 10

		k := begin
		for i, v := range bucket {
			for _, d := range v {
				array[k] = d
				k++
			}
			bucket[i] = bucket[i][:0]
		}
	}
}

func getMaxNumberOfDigits(array IntSlice, begin, end int) int {
	maxNumber := minInt
	temp := 0
	for i := begin; i <= end; i++ {
		temp = int(math.Log10(float64(array[i])) + 1)
		if temp > maxNumber {
			maxNumber = temp
		}
	}
	fmt.Println(maxNumber)
	return maxNumber
}

func getMaxNumber(array IntSlice, begin, end int) int {
	maxNumber := minInt
	for i := begin; i <= end; i++ {
		if array[i] > maxNumber {
			maxNumber = array[i]
		}
	}
	return maxNumber
}

func getDigit(integer, divisor int) int {
	return (integer / divisor) % 10
}
