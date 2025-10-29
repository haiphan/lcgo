package problems

import (
	"reflect"
	"testing"
)

func TestSmallestNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    5,
			want: 7,
		},
		{
			name: "Example 2",
			n:    1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestNumber(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
