package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacciSearchSingle(t *testing.T) {
	assert.Equal(t, FibonacciSearch([]int{1, 10, 20, 30, 40, 50, 60}, 61), -1)
	assert.Equal(t, FibonacciSearch([]int{10, 22, 35, 40, 45, 50, 80, 82, 85, 90, 100, 235}, 235), 11)
}

func TestFibonacciSearch(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test0",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 10},
			want: 1,
		},
		{
			name: "test1",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 1},
			want: 0,
		},
		{
			name: "test2",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 50},
			want: 5,
		},
		{
			name: "test3",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 60},
			want: 6,
		},
		{
			name: "test4",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 61},
			want: -1,
		},
		{
			name: "test5",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 2},
			want: -1,
		},
		{
			name: "test6",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 59},
			want: -1,
		},
		{
			name: "test7",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: -1},
			want: -1,
		},
		{
			name: "test8",
			args: args{array: []int{}, target: -1},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciSearch(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("FibonacciSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciRecursion(t *testing.T) {
	assert.Equal(t, fibonacciWithRecursive(0), 0)
	assert.Equal(t, fibonacciWithRecursive(1), 1)
	assert.Equal(t, fibonacciWithRecursive(2), 1)
	assert.Equal(t, fibonacciWithRecursive(3), 2)
	assert.Equal(t, fibonacciWithRecursive(4), 3)
	assert.Equal(t, fibonacciWithRecursive(5), 5)
	assert.Equal(t, fibonacciWithRecursive(6), 8)
	assert.Equal(t, fibonacciWithRecursive(7), 13)
	assert.Equal(t, fibonacciWithRecursive(8), 21)
	assert.Equal(t, fibonacciWithRecursive(9), 34)
	assert.Equal(t, fibonacciWithRecursive(10), 55)
	assert.Equal(t, fibonacciWithRecursive(11), 89)
	assert.Equal(t, fibonacciWithRecursive(12), 144)
	assert.Equal(t, fibonacciWithRecursive(13), 233)
	assert.Equal(t, fibonacciWithRecursive(14), 377)
}

func TestFibonacci(t *testing.T) {
	assert.Equal(t, fibonacci(0), 0)
	assert.Equal(t, fibonacci(1), 1)
	assert.Equal(t, fibonacci(2), 1)
	assert.Equal(t, fibonacci(3), 2)
	assert.Equal(t, fibonacci(4), 3)
	assert.Equal(t, fibonacci(5), 5)
	assert.Equal(t, fibonacci(6), 8)
	assert.Equal(t, fibonacci(7), 13)
	assert.Equal(t, fibonacci(8), 21)
	assert.Equal(t, fibonacci(9), 34)
	assert.Equal(t, fibonacci(10), 55)
	assert.Equal(t, fibonacci(11), 89)
	assert.Equal(t, fibonacci(12), 144)
	assert.Equal(t, fibonacci(13), 233)
	assert.Equal(t, fibonacci(14), 377)
}
