package problems

import (
	"reflect"
	"testing"
)

func TestMinBitFlips(t *testing.T) {
	tests := []struct {
		name              string
		start, goal, want int
	}{
		{
			name:  "Example 1",
			start: 10,
			goal:  7,
			want:  3,
		},
		{
			name:  "Example 2",
			start: 3,
			goal:  4,
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minBitFlips(tt.start, tt.goal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
