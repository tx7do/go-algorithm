package sorting

import "math"

// ShellSort 希尔排序
func ShellSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	i, j := 0, 0
	current := 0
	for gap := int(math.Floor(float64(length >> 1))); gap > 0; gap = int(math.Floor(float64(gap >> 1))) {
		for i = gap; i < length; i++ {
			j = i
			current = array[i]
			for j-gap >= 0 && current < array[j-gap] {
				array[j] = array[j-gap]
				j = j - gap
			}
			array[j] = current
		}
	}
}
