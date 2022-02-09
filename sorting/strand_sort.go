package sorting

// StrandSort 链排序
// @see https://en.wikipedia.org/wiki/Strand_sort
// @see https://www.geeksforgeeks.org/strand-sort/
// @see https://rosettacode.org/wiki/Sorting_algorithms/Strand_sort
func StrandSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	unsorted := make(IntSlice, length)
	sorted := make(IntSlice, 0)
	copy(unsorted, array[begin:end+1])

	for unsorted.Len() > 0 {
		subList := make(IntSlice, 0)
		subList = append(subList, unsorted[0])
		unsorted = unsorted[1:]

		for i := 0; i < unsorted.Len(); {
			if unsorted[i] > subList[subList.Len()-1] {
				subList = append(subList, unsorted[i])
				unsorted = append(unsorted[:i], unsorted[i+1:]...)

			} else {
				i++
			}
		}

		sorted = strandMerge(sorted, subList)
	}

	for i := begin; i <= end; i++ {
		array[i] = sorted[i]
	}
}

func strandMerge(left, right IntSlice) IntSlice {
	lengthLeft := left.Len()
	lengthRight := right.Len()

	var result = make(IntSlice, 0, lengthLeft+lengthRight)

	i, j := 0, 0
	for i < lengthLeft && j < lengthRight {
		if less(left[i], right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for ; i < lengthLeft; i++ {
		result = append(result, left[i])
	}
	for ; j < lengthRight; j++ {
		result = append(result, right[j])
	}

	return result
}
