package problems

import (
	"reflect"
	"testing"
)

func TestSortVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "lEetcOde",
			want: "lEOtcede",
		},
		{
			name: "Example 2",
			s:    "lYmpH",
			want: "lYmpH",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortVowels(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
