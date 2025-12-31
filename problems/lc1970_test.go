package problems

import (
	"reflect"
	"testing"
)

func TestLatestDayToCross(t *testing.T) {
	tests := []struct {
		name     string
		row, col int
		cells    [][]int
		want     int
	}{
		{
			name:  "Example 1",
			row:   2,
			col:   2,
			cells: [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}},
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := latestDayToCross(tt.row, tt.col, tt.cells)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("latestDayToCross() = %v, want %v", got, tt.want)
			}
		})
	}
}
