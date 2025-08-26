package problems

import (
	"reflect"
	"testing"
)

func TestAreaOfMaxDiagonal(t *testing.T) {
	tests := []struct {
		name       string
		dimensions [][]int
		want       int
	}{
		{
			name:       "Example 1",
			dimensions: [][]int{{9, 3}, {8, 6}},
			want:       48,
		},
		{
			name:       "Example 2",
			dimensions: [][]int{{3, 4}, {4, 3}},
			want:       12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := areaOfMaxDiagonal(tt.dimensions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("areaOfMaxDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
