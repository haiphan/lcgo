package problems

import (
	"reflect"
	"testing"
)

func TestGridGame(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int64
	}{
		{
			name: "Example 1",
			grid: [][]int{{2, 5, 4}, {1, 5, 1}},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gridGame(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gridGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
