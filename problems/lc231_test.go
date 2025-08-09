package problems

import (
	"reflect"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "Example 1",
			n:    1,
			want: true,
		},
		{
			name: "Example 2",
			n:    3,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPowerOfTwo(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
