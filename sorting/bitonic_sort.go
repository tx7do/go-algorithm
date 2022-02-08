package sorting

import (
	"math"
)

// BitonicSort 双调排序
// @see https://en.wikipedia.org/wiki/Bitonic_sorter
// @see https://kalkicode.com/bitonic-sort
// @see http://farazdagi.com/blog/2015/bitonic-sort-in-go/
// @see https://www.tutorialspoint.com/bitonic-sort-in-cplusplus
// @see https://www.geeksforgeeks.org/cpp-program-for-bitonic-sort/
// @see https://www.inf.hs-flensburg.de/lang/algorithmen/sortieren/bitonic/bitonicen.htm
// @see https://blog.csdn.net/ljiabin/article/details/8630627
func BitonicSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	var b bitonic
	sentinel := make(chan struct{})
	go b.sort(array, begin, end+1, AscendingOrder, sentinel)
	<-sentinel
}

type bitonic struct {
}

func (b bitonic) sort(array Interface, low, cnt int, dir bool, sentinel chan struct{}) {
	if cnt < 2 {
		sentinel <- struct{}{}
		return
	}

	mid := b.divisionByTwo(cnt)

	c1 := make(chan struct{})
	c2 := make(chan struct{})
	go b.sort(array, low, mid, !dir, c1)
	go b.sort(array, low+mid, cnt-mid, dir, c2)
	<-c1
	<-c2

	b.merge(array, low, cnt, dir, sentinel)
}

func (b bitonic) merge(array Interface, low, cnt int, dir bool, sentinel chan struct{}) {
	if cnt < 2 {
		sentinel <- struct{}{}
		return
	}

	mid := b.greatestPowerOfTwoLessThan(cnt)

	for i := low; i < low+cnt-mid; i++ {
		b.compareAndSwap(array, i, i+mid, dir)
	}

	c1 := make(chan struct{})
	c2 := make(chan struct{})
	go b.merge(array, low, mid, dir, c1)
	go b.merge(array, low+mid, cnt-mid, dir, c2)
	<-c1
	<-c2

	sentinel <- struct{}{}
}

func (b bitonic) compareAndSwap(array Interface, i int, j int, dir bool) {
	if array.Less(j, i) == dir {
		array.Swap(i, j)
	}
}

func (b bitonic) divisionByTwo(n int) int {
	return int(math.Floor(float64(n >> 1)))
}

func (b bitonic) greatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k < n {
		k = multiplyByTwo(k)
	}
	return divisionByTwo(k)
}
