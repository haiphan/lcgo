package problems

import (
	"reflect"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "hello",
			want: "holle",
		},
		{
			name: "Example 2",
			s:    "IceCreAm",
			want: "AceCreIm",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseVowels(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
