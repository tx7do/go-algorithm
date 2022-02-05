package sorting

import (
	"math"
)

// HeapSort 堆排序
// @see https://brilliant.org/wiki/heap-sort/
func HeapSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	var maxHeap heap
	maxHeap.buildHeap(array, begin, end)
	maxHeap.popTop(array, begin, end)
}

type heap struct {
}

func (h *heap) siftDown(data Interface, lo, hi, first int) {
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

func (h *heap) buildHeap(array Interface, begin, end int) {
	low, high := begin, end
	mid := int(math.Floor(float64(end + 1>>1)))
	for i := mid; i >= begin; i-- {
		h.siftDown(array, i, high, low)
	}
}

func (h *heap) popTop(array Interface, begin, end int) {
	first := begin
	low, high := begin, end
	for i := high; i >= begin; i-- {
		array.Swap(first, first+i)
		h.siftDown(array, low, i, first)
	}
}
