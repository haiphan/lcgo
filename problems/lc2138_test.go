package problems

import (
	"reflect"
	"testing"
)

func TestDivideString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		fill byte
		want []string
	}{
		{
			name: "Example 1",
			s:    "abcdefghi",
			k:    3,
			fill: 'x',
			want: []string{"abc", "def", "ghi"},
		},
		{
			name: "Example 2",
			s:    "abcdefghij",
			k:    3,
			fill: 'x',
			want: []string{"abc", "def", "ghi", "jxx"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divideString(tt.s, tt.k, tt.fill)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideString() = %v, want %v", got, tt.want)
			}
		})
	}
}
