package problems

import (
	"reflect"
	"testing"
)

func TestNumOfWays(t *testing.T) {
	tests := []struct {
		name    string
		n, want int
	}{
		{
			name: "Example 1",
			n:    1,
			want: 12,
		},
		{
			name: "Example 2",
			n:    5000,
			want: 30228214,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numOfWays(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
