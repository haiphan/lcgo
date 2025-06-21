package problems

import (
	"reflect"
	"testing"
)

func TestMinimumDeletions(t *testing.T) {
	tests := []struct {
		name string
		word string
		k    int
		want int
	}{
		{
			name: "Example 1",
			word: "aabcaba",
			k:    0,
			want: 3,
		},
		{
			name: "Example 2",
			word: "dabdcbdcdcd",
			k:    2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumDeletions(tt.word, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
