package search

import algorithm "go-algorithm"

func ExponentialSearch(array IntSlice, target int) int {
	length := array.Len()
	if length == 0 {
		return -1
	}

	if array[0] == target {
		return 0
	}

	if array[length-1] == target {
		return length - 1
	}

	searchRange := 1
	for searchRange < length && array[searchRange] <= target {
		searchRange = multiplyByTwo(searchRange)
	}

	startIndex := divisionByTwo(searchRange)
	endIndex := algorithm.Min(searchRange, length)
	bi := BinarySearch(array[startIndex:endIndex], target)
	if bi == -1 {
		return -1
	} else {
		return bi + startIndex
	}
}
