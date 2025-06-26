package problems

import (
	"reflect"
	"testing"
)

func TestLongestSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "Example 1",
			s:    "1001010",
			k:    5,
			want: 5,
		},
		{
			name: "Example 2",
			s:    "00101001",
			k:    1,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubsequence(tt.s, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
