package problems

import (
	"reflect"
	"testing"
)

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		chars string
		want  int
	}{
		{
			name:  "Example 1",
			words: []string{"cat", "bt", "hat", "tree"},
			chars: "attach",
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countCharacters(tt.words, tt.chars)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
