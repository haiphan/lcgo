package problems

import (
	"reflect"
	"testing"
)

func TestCountTriples(t *testing.T) {
	tests := []struct {
		name    string
		n, want int
	}{
		{
			name: "Example 1",
			n:    5,
			want: 2,
		},
		{
			name: "Example 2",
			n:    10,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTriples(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countTriples() = %v, want %v", got, tt.want)
			}
		})
	}
}
