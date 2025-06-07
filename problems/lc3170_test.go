package problems

import (
	"reflect"
	"testing"
)

func TestClearStars(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "abc",
			want: "abc",
		},
		{
			name: "Example 2",
			s:    "aaba*",
			want: "aab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clearStars(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
