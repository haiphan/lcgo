package problems

import (
	"reflect"
	"testing"
)

func TestKnightDialer(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    1,
			want: 10,
		},
		{
			name: "Example 2",
			n:    2,
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := knightDialer(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("knightDialer() = %v, want %v", got, tt.want)
			}
		})
	}
}
