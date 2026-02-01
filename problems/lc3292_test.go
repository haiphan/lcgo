package problems

import (
	"reflect"
	"testing"
)

func TestMinValidStrings(t *testing.T) {
	tests := []struct {
		name   string
		words  []string
		target string
		want   int
	}{
		{
			name:   "Example 1",
			words:  []string{"abc", "aaaaa", "bcdef"},
			target: "aabcdabc",
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minValidStrings(tt.words, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minValidStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
