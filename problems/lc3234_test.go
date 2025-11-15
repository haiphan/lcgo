package problems

import (
	"reflect"
	"testing"
)

func TestNumberOfSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "00011",
			want: 5,
		},
		{
			name: "Example 2",
			s:    "101101",
			want: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfSubstrings(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
