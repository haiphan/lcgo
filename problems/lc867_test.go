package problems

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "Example 1",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transpose(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
