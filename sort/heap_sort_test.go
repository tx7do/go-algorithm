package sort

import "testing"

func TestHeapSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort-1",
			args: args{array: []int{-100, -99, -98, 0, 98, 99, 100, 0, 1, -100, 98}},
			want: []int{-100, -100, -99, -98, 0, 0, 1, 98, 98, 99, 100},
		},
		{
			name: "sort-2",
			args: args{array: []int{5, 6, 7, 2, 1, 0}},
			want: []int{0, 1, 2, 5, 6, 7},
		},
		{
			name: "sort-3",
			args: args{array: []int{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}},
			want: []int{-590, -482, -412, -392, -381, -317, -312, -243, -215, -46, -14, 22, 158, 177, 217, 273, 380, 424, 514, 925},
		},
		{
			name: "sort-4",
			args: args{array: []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}},
			want: []int{1, 2, 3, 4, 10, 14, 16, 18, 20, 23, 24, 40, 41, 43, 46, 51, 53, 55, 56, 57, 59, 62, 68, 73, 74, 78, 79, 84, 87, 88, 89, 90, 93, 95, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.array)
			confirmed := true
			for i := 0; i < len(tt.args.array); i++ {
				if tt.args.array[i] != tt.want[i] {
					confirmed = false
				}
			}
			if !confirmed {
				t.Errorf("HeapSort() = %+v,\n want %+v", tt.args.array, tt.want)
			}
		})
	}
}
