package problems

import (
	"reflect"
	"testing"
)

func TestMaxTotalFruits(t *testing.T) {
	tests := []struct {
		name     string
		fruits   [][]int
		startPos int
		k        int
		want     int
	}{
		{
			name:     "Example 1",
			fruits:   [][]int{{2, 8}, {6, 3}, {8, 6}},
			startPos: 5,
			k:        4,
			want:     9,
		},
		{
			name:     "Example 2",
			fruits:   [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}},
			startPos: 5,
			k:        4,
			want:     14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTotalFruits(tt.fruits, tt.startPos, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxTotalFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
