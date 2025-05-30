package problems

import (
	"reflect"
	"testing"
)

func TestStoneGame(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		want  bool
	}{
		{
			name:  "Example 1",
			piles: []int{5, 3, 4, 5},
			want:  true,
		},
		{
			name:  "Example 2",
			piles: []int{3, 7, 2, 3},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stoneGame(tt.piles)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stoneGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
