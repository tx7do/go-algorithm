package search

// BinarySearch 二分查找
func BinarySearch(array []int, target int) int {
	if len(array) == 0 {
		return -1
	}

	low, high := 0, len(array)-1
	mid := 0
	for low <= high {
		//mid = low + (high-low)/2
		mid = low + (high-low)>>1
		if array[mid] > target {
			high = mid - 1
		} else if array[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// LowerBound 二分查找第一个元素的位置
func LowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		//mid = low + (high-low)/2
		mid = low + (high-low)>>1
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// UpperBound 二分查找第一个大于该元素的位置
func UpperBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		//mid = low + (high-low)/2
		mid = low + (high-low)>>1
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
