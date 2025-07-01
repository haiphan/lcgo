package problems

import (
	"reflect"
	"testing"
)

func TestPossibleStringCount(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{
			name: "Example 1",
			word: "abbcccc",
			want: 5,
		},
		{
			name: "Example 2",
			word: "abcd",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := possibleStringCount(tt.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("possibleStringCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
