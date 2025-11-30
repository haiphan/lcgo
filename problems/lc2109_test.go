package problems

import (
	"reflect"
	"testing"
)

func TestAddSpaces(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		spaces []int
		want   string
	}{
		{
			name:   "Example 1",
			s:      "LeetcodeHelpsMeLearn",
			spaces: []int{8, 13, 15},
			want:   "Leetcode Helps Me Learn",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addSpaces(tt.s, tt.spaces)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
