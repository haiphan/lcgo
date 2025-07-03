package problems

import (
	"reflect"
	"testing"
)

func TestPossibleStringCount2(t *testing.T) {
	tests := []struct {
		name string
		word string
		k    int
		want int
	}{
		{
			name: "Example 1",
			word: "aabbccdd",
			k:    7,
			want: 5,
		},
		{
			name: "Example 2",
			word: "aabbccdd",
			k:    8,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := possibleStringCount2(tt.word, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("possibleStringCount2() = %v, want %v", got, tt.want)
			}
		})
	}
}
