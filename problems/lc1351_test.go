package problems

import (
	"reflect"
	"testing"
)

func TestCountNegatives(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countNegatives(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
