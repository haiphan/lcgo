package problems

import (
	"reflect"
	"testing"
)

func TestRemoveAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{
			name:  "Example 1",
			words: []string{"vqelbi", "lbiqev", "bgbmhmnqh", "hbnmmqhbg", "mqhgnmbbh", "mqbhnbhgm", "bmgbhmnhq"},
			want:  []string{"vqelbi", "bgbmhmnqh"},
		},
		{
			name:  "Example 2",
			words: []string{"abba", "baba", "bbaa", "cd", "cd"},
			want:  []string{"abba", "cd"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeAnagrams(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
