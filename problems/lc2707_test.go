package problems

import (
	"reflect"
	"testing"
)

func TestMinExtraChar(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		dictionary []string
		want       int
	}{
		{
			name:       "Example 1",
			s:          "leetscode",
			dictionary: []string{"leet", "code", "leetcode"},
			want:       1,
		},
		{
			name:       "Example 2",
			s:          "abc",
			dictionary: []string{"a", "b", "c"},
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minExtraChar(tt.s, tt.dictionary)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minExtraChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
