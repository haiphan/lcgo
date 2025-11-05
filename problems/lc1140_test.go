package problems

import (
	"reflect"
	"testing"
)

func TestStoneGameII(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		want  int
	}{
		{
			name:  "Example 1",
			piles: []int{2, 7, 9, 4, 4},
			want:  10,
		},
		{
			name:  "Example 2",
			piles: []int{1, 2, 3, 4, 5, 100},
			want:  104,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stoneGameII(tt.piles)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stoneGameII() = %v, want %v", got, tt.want)
			}
		})
	}
}
