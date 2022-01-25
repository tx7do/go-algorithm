package search

// InterpolationSearch 插值查找
func InterpolationSearch(array []int, target int) int {
	if len(array) == 0 {
		return -1
	}

	low, high := 0, len(array)-1

	var mid = 0
	for array[low] < target && array[high] > target {
		//mid = low + (high-low)/2
		mid = low + (high-low)>>1

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
