package problems

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
