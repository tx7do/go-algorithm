package sorting

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestBogoSort(t *testing.T) {
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
			args: args{array: IntSlice{6, 8, 9, 4, 12, 1}},
			want: IntSlice{1, 4, 6, 8, 9, 12},
		},
		{
			name: "sort-3",
			args: args{array: IntSlice{2, 3, 1, 4}},
			want: IntSlice{1, 2, 3, 4},
		},
		{
			name: "sort-4",
			args: args{array: IntSlice{2, 4, 5, 3, 1}},
			want: IntSlice{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BogoSort(tt.args.array, 0, tt.args.array.Len()-1)
			confirmed := true
			for i := 0; i < tt.args.array.Len(); i++ {
				if tt.args.array[i] != tt.want[i] {
					confirmed = false
				}
			}
			if !confirmed {
				t.Errorf("BogoSort() = %+v,\n want %+v", tt.args.array, tt.want)
			}
		})
	}
}
