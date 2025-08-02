package problems

import (
	"reflect"
	"testing"
)

func TestMinCost(t *testing.T) {
	tests := []struct {
		name    string
		basket1 []int
		basket2 []int
		want    int64
	}{
		{
			name:    "Example 1",
			basket1: []int{4, 2, 2, 2},
			basket2: []int{1, 4, 1, 2},
			want:    1,
		},
		{
			name:    "Example 2",
			basket1: []int{2, 3, 4, 1},
			basket2: []int{3, 2, 5, 1},
			want:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCost(tt.basket1, tt.basket2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
