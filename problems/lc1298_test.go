package problems

import (
	"reflect"
	"testing"
)

func TestMaxCandies(t *testing.T) {
	tests := []struct {
		name           string
		status         []int
		candies        []int
		keys           [][]int
		containedBoxes [][]int
		initialBoxes   []int
		want           int
	}{
		{
			name:           "Example 1",
			status:         []int{1, 0, 1, 0},
			candies:        []int{7, 5, 4, 100},
			keys:           [][]int{{}, {}, {1}, {}},
			containedBoxes: [][]int{{1, 2}, {3}, {}, {}},
			initialBoxes:   []int{0},
			want:           16,
		},
		{
			name:           "Example 2",
			status:         []int{1, 0, 0, 0, 0, 0},
			candies:        []int{1, 1, 1, 1, 1, 1},
			keys:           [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
			containedBoxes: [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
			initialBoxes:   []int{0},
			want:           6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxCandies(tt.status, tt.candies, tt.keys, tt.containedBoxes, tt.initialBoxes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
