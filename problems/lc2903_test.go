package problems

import (
	"reflect"
	"testing"
)

func TestFindIndices(t *testing.T) {
	tests := []struct {
		name            string
		nums            []int
		indexDifference int
		valueDifference int
		want            []int
	}{
		{
			name:            "Example 1",
			nums:            []int{5, 1, 4, 1},
			indexDifference: 2,
			valueDifference: 4,
			want:            []int{0, 3},
		},
		{
			name:            "Example 2",
			nums:            []int{2, 1},
			indexDifference: 0,
			valueDifference: 0,
			want:            []int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findIndices(tt.nums, tt.indexDifference, tt.valueDifference)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
