package problems

import (
	"reflect"
	"testing"
)

func TestIsOneBitCharacter(t *testing.T) {
	tests := []struct {
		name string
		bits []int
		want bool
	}{
		{
			name: "Example 1",
			bits: []int{1, 0, 0},
			want: true,
		},
		{
			name: "Example 2",
			bits: []int{1, 1, 1, 0},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isOneBitCharacter(tt.bits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
