package problems

import (
	"reflect"
	"testing"
)

func TestIsPowerOfFourTwoSum(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "Example 1",
			n:    4,
			want: true,
		},
		{
			name: "Example 2",
			n:    5,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPowerOfFour(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
