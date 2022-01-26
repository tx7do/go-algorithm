package search

import algorithm "go-algorithm"

// FibonacciSearch 斐波那查找
func FibonacciSearch(array IntSlice, target int) int {
	length := array.Len()
	if length == 0 {
		return -1
	}

	high := length - 1
	max := array[high]

	fibMMm2 := 0              // (m-2)'th Fibonacci No.
	fibMMm1 := 1              // (m-1)'th Fibonacci No.
	fibM := fibMMm2 + fibMMm1 // m'th Fibonacci

	for fibM < max {
		fibMMm2 = fibMMm1
		fibMMm1 = fibM
		fibM = fibMMm2 + fibMMm1
	}

	var mid, offset = 0, -1
	for fibM > 1 {
		mid = algorithm.Min(offset+fibMMm2, high)

		if array[mid] < target {
			fibM = fibMMm1
			fibMMm1 = fibMMm2
			fibMMm2 = fibM - fibMMm1
			offset = mid
		} else if array[mid] > target {
			fibM = fibMMm2
			fibMMm1 = fibMMm1 - fibMMm2
			fibMMm2 = fibM - fibMMm1
		} else {
			return mid
		}
	}

	if offset < high && (array[offset+1] == target) {
		return offset + 1
	}

	return -1
}

// fibonacciWithRecursive 递归求斐波那数
// f(n) = f(n-1) + f(n-2), n >= 2
func fibonacciWithRecursive(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacciWithRecursive(n-1) + fibonacciWithRecursive(n-2)
}

// fibonacci 求斐波那数
func fibonacci(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}
