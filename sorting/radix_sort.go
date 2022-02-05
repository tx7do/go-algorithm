package sorting

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

// RadixSort 基数排序
// @see https://mp.weixin.qq.com/s/Z8gU9QLpMnA-zoMc9ZeR2w
// @see https://www.geeksforgeeks.org/radix-sort/
// @see https://en.wikipedia.org/wiki/Radix_sort
func RadixSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	maxNumber := minInt
	for i := begin; i <= end; i++ {
		if array[i] > maxNumber {
			maxNumber = array[i]
		}
	}

	n := 1
	bucket := make([]IntSlice, 10)
	for n <= maxNumber {
		for _, v := range array {
			bucket[(v/n)%10] = append(bucket[(v/n)%10], v)
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
