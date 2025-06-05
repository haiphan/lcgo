package problems

import (
	"reflect"
	"testing"
)

func TestSmallestEquivalentString(t *testing.T) {
	tests := []struct {
		name    string
		s1      string
		s2      string
		baseStr string
		want    string
	}{
		{
			name:    "Example 1",
			s1:      "parker",
			s2:      "morris",
			baseStr: "parser",
			want:    "makkek",
		},
		{
			name:    "Example 2",
			s1:      "hello",
			s2:      "world",
			baseStr: "hold",
			want:    "hdld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestEquivalentString(tt.s1, tt.s2, tt.baseStr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestEquivalentString() = %v, want %v", got, tt.want)
			}
		})
	}
}
