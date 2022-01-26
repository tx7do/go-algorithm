package sorting

// RadixSort 基数排序
// @see https://mp.weixin.qq.com/s/Z8gU9QLpMnA-zoMc9ZeR2w
func RadixSort(array IntSlice) {
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
}
