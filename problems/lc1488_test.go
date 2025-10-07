package problems

import (
	"reflect"
	"testing"
)

func TestAvoidFlood(t *testing.T) {
	tests := []struct {
		name  string
		rains []int
		want  []int
	}{
		{
			name:  "Example 1",
			rains: []int{1, 2, 3, 4},
			want:  []int{-1, -1, -1, -1},
		},
		{
			name:  "Example 2",
			rains: []int{1, 2, 0, 0, 2, 1},
			want:  []int{-1, -1, 2, 1, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := avoidFlood(tt.rains)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avoidFlood() = %v, want %v", got, tt.want)
			}
		})
	}
}
