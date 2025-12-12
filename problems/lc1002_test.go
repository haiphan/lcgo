package problems

import (
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{
			name:  "Example 1",
			words: []string{"bella", "label", "roller"},
			want:  []string{"e", "l", "l"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := commonChars(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
