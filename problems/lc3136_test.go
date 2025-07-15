package problems

import (
	"reflect"
	"testing"
)

func TestIsValidString(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "Example 1",
			word: "234Adas",
			want: true,
		},
		{
			name: "Example 2",
			word: "b3",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidString(tt.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
