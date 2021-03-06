package search

import (
	algorithm "go-algorithm"
	"math"
)

// JumpSearch 分块查找/跳点查找/跳跃查找
func JumpSearch(array IntSlice, target int) int {
	length := array.Len()
	if length == 0 {
		return -1
	}

	step := int(math.Sqrt(float64(length)))

	prev := 0
	for array[algorithm.Min(step, length)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(length)))
		if prev >= length {
			return -1
		}
	}

	for array[prev] < target {
		prev++
		if prev == algorithm.Min(step, length) {
			return -1
		}
	}

	if array[prev] == target {
		return prev
	}

	return -1
}
