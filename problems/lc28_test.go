package problems

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		name             string
		haystack, needle string
		want             int
	}{
		{
			name:     "Example 1",
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strStr(tt.haystack, tt.needle)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
