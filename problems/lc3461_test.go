package problems

import (
	"reflect"
	"testing"
)

func TestHasSameDigits(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "3902",
			want: true,
		},
		{
			name: "Example 2",
			s:    "34789",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasSameDigits(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hasSameDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
