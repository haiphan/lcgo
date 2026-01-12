package problems

import (
	"reflect"
	"testing"
)

func TestMinTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "Example 1",
			points: [][]int{{1, 1}, {3, 4}, {-1, 0}},
			want:   7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minTimeToVisitAllPoints(tt.points)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minTimeToVisitAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
