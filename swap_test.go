package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	a, b := 3, 6
	Swap(&a, &b)
	assert.Equal(t, a, 6)
	assert.Equal(t, b, 3)
}

func TestSwapXor(t *testing.T) {
	a, b := 3, 6
	SwapXor(&a, &b)
	assert.Equal(t, a, 6)
	assert.Equal(t, b, 3)
}

func TestSwapMath(t *testing.T) {
	a, b := 3, 6
	SwapMath(&a, &b)
	assert.Equal(t, a, 6)
	assert.Equal(t, b, 3)
}

func TestSwapTemp(t *testing.T) {
	a, b := 3, 6
	SwapTemp(&a, &b)
	assert.Equal(t, a, 6)
	assert.Equal(t, b, 3)
}

func TestSwapReturn(t *testing.T) {
	a, b := 3, 6
	a, b = SwapReturn(a, b)
	assert.Equal(t, a, 6)
	assert.Equal(t, b, 3)
}

func TestSwapRange(t *testing.T) {
	arr := []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	SwapRange(arr, 0, 5, 5)
	for i := 0; i < 10; i++ {
		assert.Equal(t, arr[i], i+1)
	}
}
