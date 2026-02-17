package problems

import (
	"reflect"
	"testing"
)

func TestLongestBalanced(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "abbac",
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestBalanced(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
