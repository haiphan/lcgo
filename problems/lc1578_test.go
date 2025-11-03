package problems

import (
	"reflect"
	"testing"
)

func TestMinCost1578(t *testing.T) {
	tests := []struct {
		name       string
		colors     string
		neededTime []int
		want       int
	}{
		{
			name:       "Example 1",
			colors:     "abaac",
			neededTime: []int{1, 2, 3, 4, 5},
			want:       3,
		},
		{
			name:       "Example 2",
			colors:     "abc",
			neededTime: []int{1, 2, 3},
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCost1578(tt.colors, tt.neededTime)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minCost1578() = %v, want %v", got, tt.want)
			}
		})
	}
}
