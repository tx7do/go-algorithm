package search

// BinarySearch 二分查找
func BinarySearch(array IntSlice, target int) int {
	length := array.Len()
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	mid := 0
	for low <= high {
		mid = low + divisionByTwo(high-low)
		if array[mid] == target {
			return mid
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// LowerBound 二分查找第一个元素的位置
func LowerBound(array IntSlice, target int) int {
	length := array.Len()
	if length == 0 {
		return -1
	}

	low, high, mid := 0, length-1, 0
	for low <= high {
		mid = low + divisionByTwo(high-low)
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// UpperBound 二分查找第一个大于该元素的位置
func UpperBound(array IntSlice, target int) int {
	length := array.Len()
	if length == 0 {
		return -1
	}

	low, high, mid := 0, length-1, 0
	for low <= high {
		mid = low + divisionByTwo(high-low)
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
