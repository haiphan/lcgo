package problems

import (
	"reflect"
	"testing"
)

func TestLongestSubsequenceRepeatedK(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{
			name: "Example 1",
			s:    "ababc",
			k:    2,
			want: "ab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubsequenceRepeatedK(tt.s, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestSubsequenceRepeatedK() = %v, want %v", got, tt.want)
			}
		})
	}
}
