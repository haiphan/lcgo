package problems

import (
	"reflect"
	"testing"
)

func TestMaxSubarrays(t *testing.T) {
	tests := []struct {
		name             string
		n                int
		conflictingPairs [][]int
		want             int64
	}{
		{
			name:             "Example 1",
			n:                4,
			conflictingPairs: [][]int{{1, 4}, {2, 3}},
			want:             9,
		},
		{
			name:             "Example 2",
			n:                5,
			conflictingPairs: [][]int{{1, 2}, {2, 5}, {3, 5}},
			want:             12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubarrays(tt.n, tt.conflictingPairs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
