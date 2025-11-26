package problems

import (
	"reflect"
	"testing"
)

func TestNumberOfPaths(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}},
			k:    3,
			want: 2,
		},
		{
			name: "Example 1",
			grid: [][]int{{0, 0}},
			k:    5,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfPaths(tt.grid, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
