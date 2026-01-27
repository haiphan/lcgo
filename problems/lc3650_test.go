package problems

import (
	"reflect"
	"testing"
)

func TestMinCost3650(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{
			name:  "Example 1",
			n:     4,
			edges: [][]int{{0, 1, 3}, {3, 1, 1}, {2, 3, 4}, {0, 2, 2}},
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCost3650(tt.n, tt.edges)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minCost3650() = %v, want %v", got, tt.want)
			}
		})
	}
}
