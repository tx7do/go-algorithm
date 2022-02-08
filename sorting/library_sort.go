package sorting

import "math"

// LibrarySort 图书馆排序
// @see https://en.wikipedia.org/wiki/Library_sort
// @see https://zh.wikipedia.org/wiki/%E5%9B%BE%E4%B9%A6%E9%A6%86%E6%8E%92%E5%BA%8F
// @see http://www.cs.sunysb.edu/~bender/newpub/BenderFaMo06-librarysort.pdf
// @see http://zhangyu8374.iteye.com/blog/86317
// @see http://www.cnblogs.com/richselian/archive/2011/09/16/2179152.html
// @see http://algorithm.l918.net/forum.php?mod=viewthread&tid=19
// @see http://www.softpanorama.org/Algorithms/Sorting/insertion_sort.shtml
// @see https://blog.emptyq.net/a?ID=00004-ad2225d6-ea28-413c-9e76-20db26b331e3
func LibrarySort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	//e := 3
	//rebalance_local := true
	//newSize := e*(length+1) - 1
	//arr_new := make(InterfaceSlice, newSize)
}

func LibrarySort_Pos(number, count, newSize int) int {
	return number*int(math.Floor(float64((newSize+1)/(count+1)))) - 1
}
