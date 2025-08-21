package problems

import (
	"reflect"
	"testing"
)

func TestLongestContinuousSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "abacaba",
			want: 2,
		},
		{
			name: "Example 2",
			s:    "abcde",
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestContinuousSubstring(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestContinuousSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
