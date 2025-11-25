package problems

import (
	"reflect"
	"testing"
)

func TestSmallestRepunitDivByK(t *testing.T) {
	tests := []struct {
		name    string
		k, want int
	}{
		{
			name: "Example 1",
			k:    1,
			want: 1,
		},
		{
			name: "Example 2",
			k:    2,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestRepunitDivByK(tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestRepunitDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
