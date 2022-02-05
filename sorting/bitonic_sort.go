package sorting

import (
	"fmt"
	"math"
)

// BitonicSort 双调排序
// @see https://en.wikipedia.org/wiki/Bitonic_sorter
// @see https://kalkicode.com/bitonic-sort
// @see http://farazdagi.com/blog/2015/bitonic-sort-in-go/
func BitonicSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	var b bitonic
	sentinel := make(chan struct{})
	go b.sort(array.Part(begin, end+1), AscendingOrder, sentinel)
	<-sentinel
}

type bitonic struct {
}

func (b bitonic) sort(array Interface, order bool, sentinel chan struct{}) {
	if array.Len() < 2 {
		sentinel <- struct{}{}
		return
	}

	mid := int(math.Floor(float64(array.Len() >> 1)))

	c1 := make(chan struct{})
	c2 := make(chan struct{})
	leftSide := array.Part(0, mid)
	rightSide := array.Part(mid, array.Len())
	go b.sort(leftSide, AscendingOrder, c1)
	go b.sort(rightSide, DescendingOrder, c2)
	<-c1
	<-c2

	b.merge(array, order, sentinel)
}

func (b bitonic) merge(array Interface, order bool, sentinel chan struct{}) {
	if array.Len() < 2 {
		sentinel <- struct{}{}
		return
	}

	mid := int(math.Floor(float64(array.Len() >> 1)))

	for i := 0; i < mid; i++ {
		b.compareAndSwap(array, i, i+mid, order)
	}

	if mid > 1 {
		c1 := make(chan struct{})
		c2 := make(chan struct{})
		leftSide := array.Part(0, mid)
		rightSide := array.Part(mid, array.Len())
		go b.merge(leftSide, order, c1)
		go b.merge(rightSide, order, c2)
		<-c1
		<-c2
	}

	sentinel <- struct{}{}
}

func (b bitonic) compareAndSwap(array Interface, i int, j int, order bool) {
	//fmt.Println("compareAndSwap ", array.Get(i).(int), array.Get(j).(int), array.Less(j, i))
	if array.Less(j, i) == order {
		fmt.Println(array.Get(j).(int), array.Get(i).(int))
		array.Swap(i, j)
	}
}
