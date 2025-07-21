package problems

import (
	"reflect"
	"testing"
)

func TestMakeFancyString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "leeetcode",
			want: "leetcode",
		},
		{
			name: "Example 2",
			s:    "aaabaaaa",
			want: "aabaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeFancyString(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeFancyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
