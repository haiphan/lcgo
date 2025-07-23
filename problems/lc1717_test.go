package problems

import (
	"reflect"
	"testing"
)

func TestMaximumGain(t *testing.T) {
	tests := []struct {
		name string
		s    string
		x    int
		y    int
		want int
	}{
		{
			name: "Example 1",
			s:    "cdbcbbaaabab",
			x:    4,
			y:    5,
			want: 19,
		},
		{
			name: "Example 2",
			s:    "aabbaaxybbaabb",
			x:    5,
			y:    4,
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumGain(tt.s, tt.x, tt.y)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumGain() = %v, want %v", got, tt.want)
			}
		})
	}
}
