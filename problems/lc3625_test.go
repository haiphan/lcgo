package problems

import (
	"reflect"
	"testing"
)

func TestCountTrapezoids3625(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "Example 1",
			points: [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTrapezoids3625(tt.points)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countTrapezoids() = %v, want %v", got, tt.want)
			}
		})
	}
}
