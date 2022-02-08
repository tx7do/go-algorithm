package search

// InterpolationSearch 插值查找
func InterpolationSearch(array IntSlice, target int) int {
	length := array.Len()
	if length == 0 {
		return -1
	}

	low, high := 0, length-1

	var mid = 0
	for array[low] < target && array[high] > target {
		mid = low + divisionByTwo(high-low)

		if array[mid] < target {
			low = mid + 1
		} else if array[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	if array[low] == target {
		return low
	} else if array[high] == target {
		return high
	} else {
		return -1
	}
}
