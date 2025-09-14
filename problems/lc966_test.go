package problems

import (
	"reflect"
	"testing"
)

func TestSpellchecker(t *testing.T) {
	tests := []struct {
		name     string
		wordlist []string
		queries  []string
		want     []string
	}{
		{
			name:     "Example 1",
			wordlist: []string{"KiTe", "kite", "hare", "Hare"},
			queries:  []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"},
			want:     []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
		{
			name:     "Example 2",
			wordlist: []string{"yellow"},
			queries:  []string{"YellOw"},
			want:     []string{"yellow"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spellchecker(tt.wordlist, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spellchecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
