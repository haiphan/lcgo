package problems

import (
	"reflect"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{
			name: "Example 1",
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			name: "Example 2",
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcdOfStrings(tt.str1, tt.str2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
