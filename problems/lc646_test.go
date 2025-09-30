package problems

import (
	"reflect"
	"testing"
)

func TestFindLongestChain(t *testing.T) {
	tests := []struct {
		name  string
		pairs [][]int
		want  int
	}{
		{
			name:  "Example 1",
			pairs: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:  2,
		},
		{
			name:  "Example 2",
			pairs: [][]int{{1, 2}, {7, 8}, {4, 5}},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLongestChain(tt.pairs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
