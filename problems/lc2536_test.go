package problems

import (
	"reflect"
	"testing"
)

func TestRangeAddQueries(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		queries [][]int
		want    [][]int
	}{
		{
			name:    "Example 1",
			n:       3,
			queries: [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}},
			want:    [][]int{{1, 1, 0}, {1, 2, 1}, {0, 1, 1}},
		},
		{
			name:    "Example 2",
			n:       2,
			queries: [][]int{{0, 0, 1, 1}},
			want:    [][]int{{1, 1}, {1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rangeAddQueries(tt.n, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rangeAddQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
