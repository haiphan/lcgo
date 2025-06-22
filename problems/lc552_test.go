package problems

import (
	"reflect"
	"testing"
)

func TestCheckRecord(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    2,
			want: 8,
		},
		{
			name: "Example 2",
			n:    1,
			want: 3,
		},
		{
			name: "Example 3",
			n:    10101,
			want: 183236316,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkRecord(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
