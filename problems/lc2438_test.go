package problems

import (
	"reflect"
	"testing"
)

func TestProductQueries(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		queries [][]int
		want    []int
	}{
		{
			name:    "Example 1",
			n:       15,
			queries: [][]int{{0, 1}, {2, 2}, {0, 3}},
			want:    []int{2, 4, 64},
		},
		{
			name:    "Example 2",
			n:       2,
			queries: [][]int{{0, 0}},
			want:    []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productQueries(tt.n, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
