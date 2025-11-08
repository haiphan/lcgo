package problems

import (
	"reflect"
	"testing"
)

func TestMinimumOneBitOperations(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    3,
			want: 2,
		},
		{
			name: "Example 2",
			n:    6,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumOneBitOperations(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumOneBitOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
