package problems

import (
	"reflect"
	"testing"
)

func TestNumRollsToTarget(t *testing.T) {
	tests := []struct {
		name               string
		n, k, target, want int
	}{
		{
			name:   "Example 1",
			n:      1,
			k:      6,
			target: 3,
			want:   1,
		},
		{
			name:   "Example 2",
			n:      2,
			k:      6,
			target: 7,
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numRollsToTarget(tt.n, tt.k, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
