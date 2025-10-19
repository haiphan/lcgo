package problems

import (
	"reflect"
	"testing"
)

func TestFindLexSmallestString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		a, b int
		want string
	}{
		{
			name: "Example 1",
			s:    "5525",
			a:    9,
			b:    2,
			want: "2050",
		},
		{
			name: "Example 2",
			s:    "74",
			a:    5,
			b:    1,
			want: "24",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLexSmallestString(tt.s, tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLexSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
