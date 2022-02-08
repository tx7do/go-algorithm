package search

import "math"

type IntSlice []int

func (a IntSlice) Len() int {
	return len(a)
}

func (a IntSlice) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a IntSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func divisionByTwo(n int) int {
	return int(math.Floor(float64(n >> 1)))
}

func multiplyByTwo(n int) int {
	return int(math.Floor(float64(n << 1)))
}
