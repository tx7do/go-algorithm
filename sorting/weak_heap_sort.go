package sorting

import "math"

// WeakHeapSort 弱堆排序
// @see https://en.wikipedia.org/wiki/Weak_heap
// @see https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.21.1863&rep=rep1&type=pdf
// @see https://wikichi.icu/wiki/Weak_heap
// @see https://cppsecrets.com/users/72810010410511210510797974964103109971051084699111109/C00-program-for-Weak-heap.php
// @see https://www.geeksforgeeks.org/weak-heap/
func WeakHeapSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}
}

type weakHeap struct {
}

func (h *weakHeap) buildHeap(array Interface, begin, end int) {
	low, high := begin, end
	mid := int(math.Floor(float64(end + 1>>1)))
	for i := mid; i >= begin; i-- {
		h.siftDown(array, i, high, low)
	}
}

func (h *weakHeap) siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}
