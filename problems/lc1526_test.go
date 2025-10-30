package problems

import (
	"reflect"
	"testing"
)

func TestMinNumberOperations(t *testing.T) {
	tests := []struct {
		name   string
		target []int
		want   int
	}{
		{
			name:   "Example 1",
			target: []int{1, 2, 3, 2, 1},
			want:   3,
		},
		{
			name:   "Example 1",
			target: []int{3, 1, 1, 2},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minNumberOperations(tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minNumberOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
