package problems

import (
	"reflect"
	"testing"
)

func TestMaxFrequency(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		k             int
		numOperations int
		want          int
	}{
		{
			name:          "Example 1",
			nums:          []int{1, 4, 5},
			k:             1,
			numOperations: 2,
			want:          2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxFrequency(tt.nums, tt.k, tt.numOperations)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
