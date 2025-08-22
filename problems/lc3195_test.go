package problems

import (
	"reflect"
	"testing"
)

func TestMinimumArea(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{0, 1, 0}, {1, 0, 1}},
			want: 6,
		},
		{
			name: "Example 2",
			grid: [][]int{{1, 0}, {0, 0}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumArea(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
