package sorting

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestCartesianTree(t *testing.T) {
	// 左子树的所有K1值都比该节点的K1值小，右子树则大
	{
		arr := IntSlice{5, 10, 40, 30, 28}
		tree := &cartesianTree{}
		tree.buildTree(arr, 0, arr.Len()-1)
		tree.print(os.Stdout)
		fmt.Println("")
	}
	{
		arr := IntSlice{1, 14, 5, 0, 8}
		tree := &cartesianTree{}
		tree.buildTree(arr, 0, arr.Len()-1)
		tree.print(os.Stdout)
		fmt.Println("")
	}
	{
		arr := IntSlice{9, 3, 7, 1, 8, 12, 10, 20, 15, 18, 5}
		tree := &cartesianTree{}
		tree.buildTree(arr, 0, arr.Len()-1)
		tree.print(os.Stdout)
		fmt.Println("")
	}
	{
		arr := IntSlice{11, 10, 7, 9, 5, 6, 4, 8, 2, 3, 1}
		tree := &cartesianTree{}
		tree.buildTree(arr, 0, arr.Len()-1)
		tree.print(os.Stdout)
		fmt.Println("")
	}
	{
		arr := IntSlice{7, 15, 12, 4, 10, 1, 5, 13, 6, 14, 11, 3, 9, 0, 2, 8}
		tree := &cartesianTree{}
		tree.buildTree(arr, 0, arr.Len()-1)
		tree.print(os.Stdout)
		fmt.Println("")
	}
	{
		arr := IntSlice{23, 24, 22, 30, 28, 27, 29, 21, 26, 25}
		tree := &cartesianTree{}
		tree.buildTree(arr, 0, arr.Len()-1)
		tree.print(os.Stdout)
		fmt.Println("")
	}
}

func TestCartesianTreeSort(t *testing.T) {
	type args struct {
		array IntSlice
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "sort-1",
			args: args{array: IntSlice{-100, -99, -98, 0, 98, 99, 100, 0, 1, -100, 98}},
			want: IntSlice{-100, -100, -99, -98, 0, 0, 1, 98, 98, 99, 100},
		},
		{
			name: "sort-2",
			args: args{array: IntSlice{5, 6, 7, 2, 1, 0}},
			want: IntSlice{0, 1, 2, 5, 6, 7},
		},
		{
			name: "sort-3",
			args: args{array: IntSlice{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}},
			want: IntSlice{-590, -482, -412, -392, -381, -317, -312, -243, -215, -46, -14, 22, 158, 177, 217, 273, 380, 424, 514, 925},
		},
		{
			name: "sort-4",
			args: args{array: IntSlice{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}},
			want: IntSlice{1, 2, 3, 4, 10, 14, 16, 18, 20, 23, 24, 40, 41, 43, 46, 51, 53, 55, 56, 57, 59, 62, 68, 73, 74, 78, 79, 84, 87, 88, 89, 90, 93, 95, 99},
		},
		{
			name: "sort-5",
			args: args{array: IntSlice{1, 4, 2, 5, 67, 86, 24, 63, 676, 23, 1, 3, 2, 34}},
			want: IntSlice{1, 1, 2, 2, 3, 4, 5, 23, 24, 34, 63, 67, 86, 676},
		},
		{
			name: "sort-6",
			args: args{array: IntSlice{5, 10, 40, 30, 28}},
			want: IntSlice{5, 10, 28, 30, 40},
		},
		{
			name: "sort-7",
			args: args{array: IntSlice{23, 24, 22, 30, 28, 27, 29, 21, 26, 25}},
			want: IntSlice{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		},
		{
			name: "sort-8",
			args: args{array: IntSlice{9, 3, 7, 1, 8, 12, 10, 20, 15, 18, 5}},
			want: IntSlice{1, 3, 5, 7, 8, 9, 10, 12, 15, 18, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CartesianTreeSort(tt.args.array, 0, tt.args.array.Len()-1)
			confirmed := true
			for i := 0; i < tt.args.array.Len(); i++ {
				if tt.args.array[i] != tt.want[i] {
					confirmed = false
				}
			}
			if !confirmed {
				t.Errorf("CartesianTreeSort() = %+v,\n want %+v", tt.args.array, tt.want)
			}
		})
	}
}
