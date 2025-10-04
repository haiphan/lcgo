package problems

import (
	"reflect"
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "Example 2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
