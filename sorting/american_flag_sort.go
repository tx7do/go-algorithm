package sorting

import (
	"math"
)

// AmericanFlagSort 美国旗帜排序
// @see https://en.wikipedia.org/wiki/American_flag_sort
func AmericanFlagSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	const NumberOfBuckets = 10

	maxNumber := getMaxNumber(array, begin, end)
	maxDigit := int(math.Floor(math.Log(float64(maxNumber))))

	americanFlagSort(array, begin, end, maxDigit, NumberOfBuckets)
}

func americanFlagSort(array IntSlice, begin, end int, digit, radix int) {
	offsets := make(IntSlice, radix)
	americanFlagComputeOffsets(array, begin, end, offsets, digit, radix)
	americanFlagSwap(array, begin, end, offsets, digit, radix)
	if digit == 0 {
		return
	}
	for i := 0; i < offsets.Len()-1; i++ {
		americanFlagSort(array, offsets[i], offsets[i+1]-1, digit-1, radix)
	}
}

func getRadix(x, digit, radix int) int {
	powValue := math.Pow(float64(radix), float64(digit))
	return int(math.Floor(float64(x)/powValue)) % radix
}

func americanFlagComputeOffsets(array IntSlice, begin, end int, offsets IntSlice, digit, radix int) {
	counts := make(IntSlice, radix)
	for i := begin; i <= end; i++ {
		val := getRadix(array[i], digit, radix)
		counts[val]++
	}
	sum := 0
	for i := 0; i < radix; i++ {
		offsets[i] = sum
		sum += counts[i]
	}
}

func americanFlagSwap(array IntSlice, begin, end int, offsets IntSlice, digit, radix int) {
	i := begin
	curBlock := 0
	nextFree := make(IntSlice, offsets.Len())
	copy(nextFree, offsets)
	for curBlock < radix-1 {
		if i >= begin+offsets[curBlock+1] {
			curBlock++
			continue
		}
		radixVal := getRadix(array[i], digit, radix)
		if radixVal == curBlock {
			i++
			continue
		}
		swapTo := begin + nextFree[radixVal]
		array.Swap(i, swapTo)
		nextFree[radixVal]++
	}
}
