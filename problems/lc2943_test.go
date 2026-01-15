package problems

import (
	"reflect"
	"testing"
)

func TestMaximizeSquareHoleArea(t *testing.T) {
	tests := []struct {
		name         string
		n, m         int
		hBars, vBars []int
		want         int
	}{
		{
			name:  "Example 1",
			n:     5,
			m:     4,
			hBars: []int{1, 2, 4},
			vBars: []int{1, 3},
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximizeSquareHoleArea(tt.n, tt.m, tt.hBars, tt.vBars)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximizeSquareHoleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
