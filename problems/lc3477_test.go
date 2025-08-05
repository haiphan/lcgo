package problems

import (
	"reflect"
	"testing"
)

func TestNumOfUnplacedFruits(t *testing.T) {
	tests := []struct {
		name    string
		fruits  []int
		baskets []int
		want    int
	}{
		{
			name:    "Example 1",
			fruits:  []int{4, 2, 5},
			baskets: []int{3, 5, 4},
			want:    1,
		},
		{
			name:    "Example 2",
			fruits:  []int{3, 6, 1},
			baskets: []int{6, 4, 7},
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numOfUnplacedFruits(tt.fruits, tt.baskets)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numOfUnplacedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
