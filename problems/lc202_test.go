package problems

import (
	"reflect"
	"testing"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "Example 1",
			n:    19,
			want: true,
		},
		{
			name: "Example 2",
			n:    2,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isHappy(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
