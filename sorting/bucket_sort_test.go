package sorting

import "testing"

func TestBucketSort(t *testing.T) {
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
			args: args{array: IntSlice{101, 99, 98, 0, 98, 99, 102, 0, 1, 103, 98}},
			want: IntSlice{0, 0, 1, 98, 98, 98, 99, 99, 101, 102, 103},
		},
		{
			name: "sort-2",
			args: args{array: IntSlice{5, 6, 7, 2, 1, 0}},
			want: IntSlice{0, 1, 2, 5, 6, 7},
		},
		{
			name: "sort-3",
			args: args{array: IntSlice{317, 381, 14, 215, 590, 243, 412, 380, 312, 925, 158, 46, 177, 22, 482, 273, 217, 514, 392, 424}},
			want: IntSlice{14, 22, 46, 158, 177, 215, 217, 243, 273, 312, 317, 380, 381, 392, 412, 424, 482, 514, 590, 925},
		},
		{
			name: "sort-4",
			args: args{array: IntSlice{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}},
			want: IntSlice{1, 2, 3, 4, 10, 14, 16, 18, 20, 23, 24, 40, 41, 43, 46, 51, 53, 55, 56, 57, 59, 62, 68, 73, 74, 78, 79, 84, 87, 88, 89, 90, 93, 95, 99},
		},
		{
			name: "sort-5",
			args: args{array: IntSlice{9, 5, 2, 4, 0, 8, 6, 4, 3, 1, 7, 4, 2, 5, 7, 4, 3, 2}},
			want: IntSlice{0, 1, 2, 2, 2, 3, 3, 4, 4, 4, 4, 5, 5, 6, 7, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BucketSort(tt.args.array, 0, tt.args.array.Len()-1)
			confirmed := true
			for i := 0; i < tt.args.array.Len(); i++ {
				if tt.args.array[i] != tt.want[i] {
					confirmed = false
				}
			}
			if !confirmed {
				t.Errorf("BucketSort() = %+v,\n want %+v", tt.args.array, tt.want)
			}
		})
	}
}

func TestBucketSortNegative(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BucketSort(tt.args.array, 0, tt.args.array.Len()-1)
			confirmed := true
			for i := 0; i < tt.args.array.Len(); i++ {
				if tt.args.array[i] != tt.want[i] {
					confirmed = false
				}
			}
			if !confirmed {
				t.Errorf("BucketSort() = %+v,\n want %+v", tt.args.array, tt.want)
			}
		})
	}
}
