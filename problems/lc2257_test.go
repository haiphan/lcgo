package problems

import (
	"reflect"
	"testing"
)

func TestCountUnguarded(t *testing.T) {
	tests := []struct {
		name          string
		m, n          int
		guards, walls [][]int
		want          int
	}{
		{
			name:   "Example 1",
			m:      4,
			n:      6,
			guards: [][]int{{0, 0}, {1, 1}, {2, 3}},
			walls:  [][]int{{0, 1}, {2, 2}, {1, 4}},
			want:   7,
		},
		{
			name:   "Example 2",
			m:      3,
			n:      3,
			guards: [][]int{{1, 1}},
			walls:  [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countUnguarded(tt.m, tt.n, tt.guards, tt.walls)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countUnguarded() = %v, want %v", got, tt.want)
			}
		})
	}
}
