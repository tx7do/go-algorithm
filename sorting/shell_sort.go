package sorting

import "math"

// ShellSort 希尔排序
// @see https://en.wikipedia.org/wiki/Shellsort
func ShellSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	i, j := 0, 0
	gap := int(math.Floor(float64(begin + (end-begin)>>1)))
	for ; gap > begin; gap = int(math.Floor(float64(gap >> 1))) {
		for i = gap; i <= end; i++ {
			for j = i - gap; j >= begin; j -= gap {
				if array.Less(j+gap, j) {
					array.Swap(j, j+gap)
				} else {
					break
				}
			}
		}
	}
}
