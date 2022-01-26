package sorting

import "math"

// HeapSort 堆排序
func HeapSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	mid := int(math.Floor(float64(length >> 1)))
	for i := mid; i >= 0; i-- {
		adjustHeap(array, i, length)
	}

	for j := length - 1; j > 0; j-- {
		//array[0], array[j] = array[j], array[0]
		array.Swap(0, j)
		adjustHeap(array, 0, j)
	}
}

// adjustHeap 调整堆
func adjustHeap(array IntSlice, root, length int) {
	temp := array[root]

	for k := root<<1 + 1; k < length; k = k<<1 + 1 {
		if (k+1 < length) && (array[k] < array[k+1]) {
			k++
		}
		if array[k] > temp {
			array[root] = array[k]
			root = k
		} else {
			break
		}
	}
	array[root] = temp
}
