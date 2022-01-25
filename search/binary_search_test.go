package search

import (
	"fmt"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
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
			if got := BinarySearch(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{array: []int{1, 10, 10, 10, 20, 30, 40, 50, 60}, target: 10},
			want: 4,
		},
		{
			name: "test1",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 1},
			want: 1,
		},
		{
			name: "test2",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 50},
			want: 6,
		},
		{
			name: "test3",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 60},
			want: 7,
		},
		{
			name: "test4",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 61},
			want: 7,
		},
		{
			name: "test5",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 2},
			want: 1,
		},
		{
			name: "test6",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 59},
			want: 6,
		},
		{
			name: "test7",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60, 60, 60}, target: 60},
			want: 9,
		},
		{
			name: "test8",
			args: args{array: []int{}, target: 1},
			want: 0,
		},
		{
			name: "test9",
			args: args{array: []int{-5, -4, -4, -4}, target: -5},
			want: 1,
		},
		{
			name: "test10",
			args: args{array: []int{1, 5, 5, 5, 5, 7, 7, 7, 7, 9}, target: 0},
			want: 0,
		},
		{
			name: "test11",
			args: args{array: []int{1, 2, 5, 8, 10}, target: 10},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperBound(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("UpperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerBound(t *testing.T) {
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
			args: args{array: []int{1, 10, 10, 10, 20, 30, 40, 50, 60}, target: 10},
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
			want: 7,
		},
		{
			name: "test5",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 2},
			want: 1,
		},
		{
			name: "test6",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 59},
			want: 6,
		},
		{
			name: "test7",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60, 60, 60}, target: 60},
			want: 6,
		},
		{
			name: "test8",
			args: args{array: []int{}, target: 1},
			want: 0,
		},
		{
			name: "test9",
			args: args{array: []int{-5, -2, -2, -2}, target: -4},
			want: 1,
		},
		{
			name: "test10",
			args: args{array: []int{1, 5, 5, 5, 5, 7, 7, 7, 7, 9}, target: 0},
			want: 0,
		},
		{
			name: "test11",
			args: args{array: []int{1, 2, 5, 8, 10}, target: 11},
			want: 5,
		},
		{
			name: "test12",
			args: args{array: []int{1, 2, 2, 2, 3, 3, 10}, target: 3},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerBound(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("LowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortSearchUpper(t *testing.T) {
	// 搜索按升序排序的列表
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	length := len(a)
	x := 6

	i := sort.Search(length,
		func(i int) bool {
			return a[i] >= x
		},
	)

	if i < length && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}

func TestSortSearchLower(t *testing.T) {
	// 搜索按降序排序的列表
	a := []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
	length := len(a)
	x := 6

	i := sort.Search(length,
		func(i int) bool {
			return a[i] <= x
		},
	)

	if i < length && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
