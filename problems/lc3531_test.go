package problems

import (
	"reflect"
	"testing"
)

func TestCountCoveredBuildings(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		buildings [][]int
		want      int
	}{
		{
			name:      "Example 1",
			n:         3,
			buildings: [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}},
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countCoveredBuildings(tt.n, tt.buildings)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countCoveredBuildings() = %v, want %v", got, tt.want)
			}
		})
	}
}
