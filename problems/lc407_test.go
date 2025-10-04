package problems

import (
	"reflect"
	"testing"
)

func TestTrapRainWater(t *testing.T) {
	tests := []struct {
		name      string
		heightMap [][]int
		want      int
	}{
		{
			name:      "Example 1",
			heightMap: [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}},
			want:      4,
		},
		{
			name:      "Example 2",
			heightMap: [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}},
			want:      10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trapRainWater(tt.heightMap)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trapRainWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
