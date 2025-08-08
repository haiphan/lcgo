package problems

import (
	"reflect"
	"testing"
)

func TestSoupServings(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want float64
	}{
		{
			name: "Example 1",
			n:    50,
			want: 0.625,
		},
		{
			name: "Example 2",
			n:    100,
			want: 0.71875,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := soupServings(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("soupServings() = %v, want %v", got, tt.want)
			}
		})
	}
}
