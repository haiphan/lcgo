package problems

import (
	"reflect"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		{
			name:     "Example 1",
			triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			want:     11,
		},
		{
			name:     "Example 2",
			triangle: [][]int{{-10}},
			want:     -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumTotal(tt.triangle)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
