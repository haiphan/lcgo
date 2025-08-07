package problems

import (
	"reflect"
	"testing"
)

func TestMaxCollectedFruits(t *testing.T) {
	tests := []struct {
		name   string
		fruits [][]int
		want   int
	}{
		{
			name:   "Example 1",
			fruits: [][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			want:   100,
		},
		{
			name:   "Example 2",
			fruits: [][]int{{1, 1}, {1, 1}},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxCollectedFruits(tt.fruits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxCollectedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
