package sorting

import (
	cHeap "container/heap"
	"sort"
)

// PatienceSort 耐心排序
// @see https://en.wikipedia.org/wiki/Patience_sorting
// @see https://rosettacode.org/wiki/Sorting_algorithms/Patience_sort
func PatienceSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	var piles []IntPile

	i := 0
	findPos := 0
	for i = begin; i <= end; i++ {
		findPos = sort.Search(len(piles), func(k int) bool { return piles[k].Top() >= array[i] })
		if findPos != len(piles) {
			piles[findPos] = append(piles[findPos], array[i])
		} else {
			piles = append(piles, IntPile{array[i]})
		}
	}

	hp := IntPilesHeap(piles)
	cHeap.Init(&hp)
	for i = begin; i <= end; i++ {
		smallPile := cHeap.Pop(&hp).(IntPile)
		array[i] = smallPile.Pop()
		if len(smallPile) != 0 {
			cHeap.Push(&hp, smallPile)
		}
	}
}

type IntPile []int

func (p IntPile) Top() int { return p[len(p)-1] }
func (p *IntPile) Pop() int {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

type IntPilesHeap []IntPile

func (h IntPilesHeap) Len() int            { return len(h) }
func (h IntPilesHeap) Less(i, j int) bool  { return h[i].Top() < h[j].Top() }
func (h IntPilesHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntPilesHeap) Push(x interface{}) { *h = append(*h, x.(IntPile)) }
func (h *IntPilesHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
