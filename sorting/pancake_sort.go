package sorting

// PancakeSort 煎饼排序
// @see https://en.wikipedia.org/wiki/Pancake_sorting
// @see https://rosettacode.org/wiki/Sorting_algorithms/Pancake_sort
// @see https://iq.opengenus.org/pancake-sort/
// @see https://www.geeksforgeeks.org/pancake-sorting/
func PancakeSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	maxIndex, currentIndex := begin, end
	for ; currentIndex > 0; currentIndex-- {
		maxIndex = pancakeFindMax(array, begin, currentIndex)
		if maxIndex != currentIndex {
			pancakeFlip(array, begin, maxIndex)
			pancakeFlip(array, begin, currentIndex)
		}
	}
}

func pancakeFindMax(array Interface, begin, end int) int {
	maxIndex := 0
	for i := begin; i <= end; i++ {
		if array.Less(maxIndex, i) {
			maxIndex = i
		}
	}
	return maxIndex
}

func pancakeFlip(array Interface, begin, end int) {
	startPos, endPos := begin, end
	for startPos < endPos {
		array.Swap(startPos, endPos)
		startPos++
		endPos--
	}
}
