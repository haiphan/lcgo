package problems

import (
	"reflect"
	"testing"
)

func TestIsStrictlyPalindromic(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "Example 1",
			n:    7,
			want: false,
		},
		{
			name: "Example 2",
			n:    123,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isStrictlyPalindromic(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isStrictlyPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
