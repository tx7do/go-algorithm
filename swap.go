package algorithm

// Swap 直接交换值,使用了golang的语法糖,跟python用法一样.
func Swap(a, b *int) {
	*a, *b = *b, *a
}

// SwapReturn 使用了golang的多返回值特性
func SwapReturn(a, b int) (int, int) {
	return b, a
}

// SwapXor 利用了位的异或运算
func SwapXor(a, b *int) {
	//*a = (*a) ^ (*b)
	//*b = (*a) ^ (*b)
	//*a = (*a) ^ (*b)
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

// SwapMath 利用了数学加减计算
func SwapMath(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

// SwapTemp 常规做法,创建了一个临时的中间值用于交换
func SwapTemp(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// SwapRange 区间交换
func SwapRange(data []int, a, b, n int) {
	for i := 0; i < n; i++ {
		Swap(&data[a+i], &data[b+i])
	}
}
