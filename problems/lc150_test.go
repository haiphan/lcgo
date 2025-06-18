package problems

import (
	"reflect"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "Example 1",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "Example 2",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := evalRPN(tt.tokens)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
