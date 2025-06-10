package problems

import (
	"reflect"
	"testing"
)

func TestMaxDifference(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "aaaaabbc",
			want: 3,
		},
		{
			name: "Example 2",
			s:    "abcabcab",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDifference(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
