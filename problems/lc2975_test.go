package problems

import (
	"reflect"
	"testing"
)

func TestMaximizeSquareArea(t *testing.T) {
	tests := []struct {
		name             string
		m, n             int
		hFences, vFences []int
		want             int
	}{
		{
			name:    "Example 1",
			m:       4,
			n:       3,
			hFences: []int{2, 3},
			vFences: []int{2},
			want:    4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximizeSquareArea(tt.m, tt.n, tt.hFences, tt.vFences)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximizeSquareArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
