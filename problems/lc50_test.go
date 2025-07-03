package problems

import (
	"reflect"
	"testing"
)

const EPSILON = 1e-6

func isEqual(a, b float64) bool {
	if a < b {
		return b-a < EPSILON
	}
	return a-b < EPSILON
}

func TestMyPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{
			name: "Example 1",
			x:    2.0,
			n:    10,
			want: 1024.0,
		},
		{
			name: "Example 2",
			x:    2.1,
			n:    3,
			want: 9.261,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myPow(tt.x, tt.n)
			if !reflect.DeepEqual(isEqual(tt.want, got), true) {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
