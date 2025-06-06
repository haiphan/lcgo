package problems

import (
	"reflect"
	"testing"
)

func TestRobotWithString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "cba",
			want: "abc",
		},
		{
			name: "Example 2",
			s:    "bac",
			want: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := robotWithString(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("robotWithString() = %v, want %v", got, tt.want)
			}
		})
	}
}
