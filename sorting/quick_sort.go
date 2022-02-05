package sorting

// QuickSort 快速排序
// @see https://en.wikipedia.org/wiki/Quicksort
// @see https://www.eecs.yorku.ca/course_archive/2010-11/W/2011/Notes/s4_quick_sort.pdf
func QuickSort(array Interface, begin, end int) {
	if begin < end {
		pi := quickSortPartition(array, begin, end)
		QuickSort(array, begin, pi-1)
		QuickSort(array, pi+1, end)
	}
}

func quickSortPartition(array Interface, begin, end int) int {
	pivot := array.Get(end)
	pi := begin - 1

	for i := begin; i < end; i++ {
		if Less(array.Get(i), pivot) {
			pi++
			array.Swap(pi, i)
		}
	}
	array.Swap(pi+1, end)

	return pi + 1
}
