package problems

import (
	"reflect"
	"testing"
)

func TestMaximumSafenessFactor(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}},
			want: 2,
		},
		{
			name: "Example 2",
			grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumSafenessFactor(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumSafenessFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
