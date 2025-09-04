package problems

import (
	"reflect"
	"testing"
)

func TestMinimumSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{1, 0, 1}, {1, 1, 1}},
			want: 5,
		},
		{
			name: "Example 2",
			grid: [][]int{{1, 0, 1, 0}, {0, 1, 0, 1}},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumSum(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
