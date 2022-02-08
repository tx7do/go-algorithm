package sorting

// ShellSort 希尔排序
// @see https://en.wikipedia.org/wiki/Shellsort
func ShellSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	i, j := 0, 0
	gap := divisionByTwo(begin + (end - begin))
	for ; gap > begin; gap = divisionByTwo(gap) {
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
