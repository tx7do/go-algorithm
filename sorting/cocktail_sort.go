package sorting

// CocktailSort 鸡尾酒排序
// @see https://zh.wikipedia.org/zh-hans/%E9%B8%A1%E5%B0%BE%E9%85%92%E6%8E%92%E5%BA%8F
// @see https://zhuanlan.zhihu.com/p/125008445
func CocktailSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	low, high := 0, length-1
	i := 0
	swapped := true
	for swapped {
		swapped = false

		for i = low; i < high; i++ {
			if array[i] > array[i+1] {
				//array[i], array[i+1] = array[i+1], array[i]
				array.Swap(i, i+1)
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false

		high--

		for i = high - 1; i >= low; i-- {
			if array[i] > array[i+1] {
				//array[i], array[i+1] = array[i+1], array[i]
				array.Swap(i, i+1)
				swapped = true
			}
		}

		low++
	}
}
