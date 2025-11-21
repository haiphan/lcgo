package problems

import (
	"reflect"
	"testing"
)

func TestCountPalindromicSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "aabca",
			want: 3,
		},
		{
			name: "Example 2",
			s:    "adc",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPalindromicSubsequence(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPalindromicSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
