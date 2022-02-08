package sorting

import (
	algorithm "go-algorithm"
)

// TimSort 蒂姆排序
// @see http://svn.python.org/projects/python/trunk/Objects/listsort.txt
// @see http://www.envisage-project.eu/proving-android-java-and-python-sorting-algorithm-is-broken-and-how-to-fix-it/
// @see https://v8.dev/blog/array-sort#timsort
// @see https://yuanchieh.page/posts/2019/2019-08-09_v8-%E5%85%A7%E7%9A%84%E6%8E%92%E5%BA%8F%E6%BC%94%E7%AE%97%E6%B3%95-timsort/
// @see https://www.bigocheatsheet.com/
// @see https://www.javatpoint.com/tim-sort
func TimSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	const runSize = 32

	left, mid, right := 0, 0, 0

	for i := begin; i < end; i += runSize {
		left = i
		right = algorithm.Min(i+runSize+1, end)
		InsertionSort(array, left, right)
	}

	for size := runSize; size <= end; size = multiplyByTwo(size) {
		for left = begin; left <= end; left += multiplyByTwo(size) {
			mid = left + size - 1
			right = algorithm.Min(left+multiplyByTwo(size)-1, end)

			if mid < right {
				merge(array, left, mid, right)
			}
		}
	}
}
