package sort

import algorithm "go-algorithm"

// RadixSort 基数排序
// https://mp.weixin.qq.com/s/Z8gU9QLpMnA-zoMc9ZeR2w
func RadixSort(array []int) {
	length := len(array)
	if length < 2 {
		return
	}

	maxValue := array[0]
	for _, val := range array {
		maxValue = algorithm.Max(maxValue, val)
	}
}
