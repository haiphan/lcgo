package problems

import (
	"reflect"
	"testing"
)

func TestMinCost3651(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{1, 2}, {2, 3}, {3, 4}},
			k:    1,
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCost3651(tt.grid, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minCost3651() = %v, want %v", got, tt.want)
			}
		})
	}
}
