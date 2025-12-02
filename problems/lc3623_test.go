package problems

import (
	"reflect"
	"testing"
)

func TestCountTrapezoids(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "Example 1",
			points: [][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}},
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTrapezoids(tt.points)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countTrapezoids() = %v, want %v", got, tt.want)
			}
		})
	}
}
