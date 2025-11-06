package problems

import (
	"reflect"
	"testing"
)

func TestMaxDistance1162(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
			want: 2,
		},
		{
			name: "Example 2",
			grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDistance1162(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDistance1162() = %v, want %v", got, tt.want)
			}
		})
	}
}
