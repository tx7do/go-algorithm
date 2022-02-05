package sorting

// CombSort 梳排序
// @see https://en.wikipedia.org/wiki/Flashsort
func CombSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	gap := end + 1
	swapped := false
	i := 0
	for gap > 1 || swapped {
		gap = updateGap(gap)
		swapped = false
		for i = begin; i < (end - gap + 1); i++ {
			if array.Less(i+gap, i) {
				array.Swap(i, i+gap)
				swapped = true
			}
		}
	}
}

func updateGap(gap int) int {
	if gap < 1 {
		return 1
	} else {
		return int(float64(gap) / 1.3)
	}
}
