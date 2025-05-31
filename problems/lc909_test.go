package problems

import (
	"reflect"
	"testing"
)

func TestSnakesAndLadders(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  int
	}{
		{
			name: "Example 1",
			board: [][]int{
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 35, -1, -1, 13, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 15, -1, -1, -1, -1},
			},
			want: 4,
		},
		{
			name: "Example 2",
			board: [][]int{
				{-1, -1},
				{-1, 3},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := snakesAndLadders(tt.board)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("snakesAndLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
