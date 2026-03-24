package problems

import (
	"reflect"
	"testing"
)

func TestConstructProductMatrix(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want [][]int
	}{
		{
			name: "Example 1",
			grid: [][]int{{1, 2}, {3, 4}},
			want: [][]int{{24, 12}, {8, 6}},
		},
		{
			name: "Example 2",
			grid: [][]int{{12345}, {2}, {1}},
			want: [][]int{{2}, {0}, {0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructProductMatrix(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructProductMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
