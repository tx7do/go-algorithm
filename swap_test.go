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
