package problems

import (
	"reflect"
	"testing"
)

func TestMaxStability(t *testing.T) {
	tests := []struct {
		name  string
		n, k  int
		edges [][]int
		want  int
	}{
		{
			name:  "Example 1",
			n:     3,
			k:     2,
			edges: [][]int{{0, 1, 2, 1}, {1, 2, 3, 0}},
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxStability(tt.n, tt.edges, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxStability() = %v, want %v", got, tt.want)
			}
		})
	}
}
