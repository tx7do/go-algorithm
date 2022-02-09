package sorting

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestStrandSort(t *testing.T) {
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
			args: args{array: IntSlice{5, 6, 7, 2, 1, 0}},
			want: IntSlice{0, 1, 2, 5, 6, 7},
		},
		{
			name: "sort-2",
			args: args{array: IntSlice{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}},
			want: IntSlice{1, 2, 3, 4, 10, 14, 16, 18, 20, 23, 24, 40, 41, 43, 46, 51, 53, 55, 56, 57, 59, 62, 68, 73, 74, 78, 79, 84, 87, 88, 89, 90, 93, 95, 99},
		},
		{
			name: "sort-3",
			args: args{array: IntSlice{802, 630, 20, 745, 52, 300, 612, 932, 78, 187}},
			want: IntSlice{20, 52, 78, 187, 300, 612, 630, 745, 802, 932},
		},
		{
			name: "sort-4",
			args: args{array: []int{802, 630, 20, 745, 52, 300, 612, 932, 78, 187}},
			want: []int{20, 52, 78, 187, 300, 612, 630, 745, 802, 932},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StrandSort(tt.args.array, 0, tt.args.array.Len()-1)
			confirmed := true
			for i := 0; i < tt.args.array.Len(); i++ {
				if tt.args.array[i] != tt.want[i] {
					confirmed = false
				}
			}
			if !confirmed {
				t.Errorf("StrandSort() = %+v,\n want %+v", tt.args.array, tt.want)
			}
		})
	}
}
