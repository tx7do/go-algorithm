package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterpolationSearchSingle(t *testing.T) {
	assert.Equal(t, InterpolationSearch(IntSlice{1, 10, 20, 30, 40, 50, 60}, 61), -1)
	assert.Equal(t, InterpolationSearch(IntSlice{5, 10, 20, 30, 40, 50, 60}, 1), -1)
}

func TestInterpolationSearch(t *testing.T) {
	type args struct {
		array  IntSlice
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test0",
			args: args{array: IntSlice{1, 10, 20, 30, 40, 50, 60}, target: 10},
			want: 1,
		},
		{
			name: "test1",
			args: args{array: IntSlice{1, 10, 20, 30, 40, 50, 60}, target: 1},
			want: 0,
		},
		{
			name: "test2",
			args: args{array: IntSlice{1, 10, 20, 30, 40, 50, 60}, target: 50},
			want: 5,
		},
		{
			name: "test3",
			args: args{array: IntSlice{1, 10, 20, 30, 40, 50, 60}, target: 60},
			want: 6,
		},
		{
			name: "test4",
			args: args{array: IntSlice{1, 10, 20, 30, 40, 50, 60}, target: 61},
			want: -1,
		},
		{
			name: "test5",
			args: args{array: IntSlice{1, 10, 20, 30, 40, 50, 60}, target: 2},
			want: -1,
		},
		{
			name: "test6",
			args: args{array: IntSlice{1, 10, 20, 30, 40, 50, 60}, target: 59},
			want: -1,
		},
		{
			name: "test7",
			args: args{array: IntSlice{1, 10, 20, 30, 40, 50, 60}, target: -1},
			want: -1,
		},
		{
			name: "test8",
			args: args{array: IntSlice{}, target: -1},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterpolationSearch(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("InterpolationSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
