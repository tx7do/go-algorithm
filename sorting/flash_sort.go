package sorting

// FlashSort 闪电排序
// @see https://en.wikipedia.org/wiki/Flashsort
func FlashSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	minValue, maxValue := array[begin], array[begin]
	i := 0
	for i = begin + 1; i <= end; i++ {
		if array[i] < minValue {
			minValue = array[i]
		}
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}

	if minValue == maxValue {
		return
	}

	m := int(0.43 * float64(length))
	c := float64((m - 1) / (maxValue - minValue))

	L := make(IntSlice, m+1)

	for i = begin; i <= end; i++ {
		k := int(c*float64(array[i]-minValue)) + 1
		L[k]++
	}
	for i = 2; i <= m; i++ {
		L[i] += L[i-1]
	}

	flash, move, j, k := 0, 0, 0, m
	for move < end {
		for j >= L[k] {
			j++
			k = int(c*float64(array[j]-minValue) + 1)
		}
		flash = array[j]
		for j < L[k] {
			k = int(c*float64(flash-minValue) + 1)
			L[k]--
			flash, array[L[k]] = array[L[k]], flash
			move++
		}
	}

	key := 0
	for k = 1; k < m; k++ {
		for i = L[k] + 1; i < L[k+1]; i++ {
			key = array[i]
			for j = i - 1; j >= 0 && array[j] > key; {
				array.Swap(j, j+1)
				j--
			}
			array[j+1] = key
		}
	}
}
