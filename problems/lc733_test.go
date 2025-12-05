package problems

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		sr, sc int
		color  int
		want   [][]int
	}{
		{
			name:  "Example 1",
			image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:    1,
			sc:    1,
			color: 2,
			want:  [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := floodFill(tt.image, tt.sr, tt.sc, tt.color)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
