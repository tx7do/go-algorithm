package sorting

// CycleSort 圈排序
// @see https://en.wikipedia.org/wiki/Cycle_sort
func CycleSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	item, pos, i := 0, 0, 0
	for cycleStart := begin; cycleStart <= end-1; cycleStart++ {
		item = array[cycleStart]

		pos = cycleStart

		for i = cycleStart + 1; i <= end; i++ {
			if array[i] < item {
				pos++
			}
		}

		if pos == cycleStart {
			continue
		}

		for item == array[pos] {
			pos += 1
		}

		if pos != cycleStart {
			item, array[pos] = array[pos], item
		}

		for pos != cycleStart {
			pos = cycleStart

			for i = cycleStart + 1; i <= end; i++ {
				if array[i] < item {
					pos += 1
				}
			}

			for item == array[pos] {
				pos += 1
			}

			if item != array[pos] {
				item, array[pos] = array[pos], item
			}
		}
	}
}
