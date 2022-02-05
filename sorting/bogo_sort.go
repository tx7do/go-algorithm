package sorting

import (
	"go-algorithm/rand"
)

// BogoSort 猴子排序
// @see https://en.wikipedia.org/wiki/Bogosort
// @see https://zh.wikipedia.org/zh-hant/Bogo%E6%8E%92%E5%BA%8F
func BogoSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	for !IsSorted(array.Part(begin, end+1)) {
		shuffle(array, begin, end)
	}
}

func shuffle(array Interface, begin, end int) {
	for i := begin; i <= end; i++ {
		j := rand.RandomInt(begin, end)
		array.Swap(i, j)
	}
}
