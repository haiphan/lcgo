package problems

import (
	"reflect"
	"testing"
)

func TestMaxSideLength(t *testing.T) {
	tests := []struct {
		name      string
		mat       [][]int
		threshold int
		want      int
	}{
		{
			name:      "Example 1",
			mat:       [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}},
			threshold: 4,
			want:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSideLength(tt.mat, tt.threshold)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSideLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
