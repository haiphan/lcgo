package problems

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "Example 1",
			x:    11,
			want: true,
		},
		{
			name: "Example 2",
			x:    12,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
