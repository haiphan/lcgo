package problems

import (
	"reflect"
	"testing"
)

func TestFindWordsContaining(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		x     byte
		want  []int
	}{
		{
			name:  "Example 1",
			words: []string{"hello", "world", "leetcode"},
			x:     'l',
			want:  []int{0, 1, 2},
		},
		{
			name:  "Example 2",
			words: []string{"abc", "def", "ghi"},
			x:     'z',
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWordsContaining(tt.words, tt.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
