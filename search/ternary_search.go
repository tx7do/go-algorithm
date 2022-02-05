package search

// TernarySearch 三叉树查找
func TernarySearch(array IntSlice, target int) int {
	length := end - begin + 1
	if length == 0 {
		return -1
	}

	low, high := 0, length-2
	mid1, mid2 := 0, 0
	for low <= high {
		mid1 = low + (high-low)/3
		mid2 = high + (high-low)/3

		if array[mid1] == target {
			return mid1
		} else if array[mid2] == target {
			return mid2
		}

		if target < array[mid1] {
			high = mid1 - 1
		} else if target > array[mid2] {
			low = mid2 + 1
		} else {
			low = mid1 + 1
			high = mid2 - 1
		}
	}

	return -1
}
