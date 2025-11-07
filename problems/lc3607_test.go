package problems

import (
	"reflect"
	"testing"
)

func TestProcessQueries(t *testing.T) {
	tests := []struct {
		name        string
		c           int
		connections [][]int
		queries     [][]int
		want        []int
	}{
		{
			name:        "Example 1",
			c:           5,
			connections: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			queries:     [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}},
			want:        []int{3, 2, 3},
		},
		{
			name:        "Example 2",
			c:           3,
			connections: [][]int{},
			queries:     [][]int{{1, 1}, {2, 1}, {1, 1}},
			want:        []int{1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processQueries(tt.c, tt.connections, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
