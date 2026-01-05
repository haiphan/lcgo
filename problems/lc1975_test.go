package problems

import (
	"reflect"
	"testing"
)

func TestMaxMatrixSum(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int64
	}{
		{
			name:   "Example 1",
			matrix: [][]int{{1, -1}, {-1, 1}},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxMatrixSum(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxMatrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
