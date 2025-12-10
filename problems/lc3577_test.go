package problems

import (
	"reflect"
	"testing"
)

func TestCountPermutations(t *testing.T) {
	tests := []struct {
		name       string
		complexity []int
		want       int
	}{
		{
			name:       "Example 1",
			complexity: []int{1, 2, 3},
			want:       2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPermutations(tt.complexity)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
