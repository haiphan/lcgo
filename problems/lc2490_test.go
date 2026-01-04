package problems

import (
	"reflect"
	"testing"
)

func TestIsCircularSentence(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		want     bool
	}{
		{
			name:     "Example 1",
			sentence: "leetcode exercises sound delightful",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isCircularSentence(tt.sentence)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isCircularSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
