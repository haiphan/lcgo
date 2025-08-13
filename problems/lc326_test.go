package problems

import (
	"reflect"
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "Example 1",
			n:    27,
			want: true,
		},
		{
			name: "Example 2",
			n:    0,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPowerOfThree(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
