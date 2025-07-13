package problems

import (
	"reflect"
	"testing"
)

func TestEarliestAndLatest(t *testing.T) {
	tests := []struct {
		name         string
		n            int
		firstPlayer  int
		secondPlayer int
		want         []int
	}{
		{
			name:         "Example 1",
			n:            11,
			firstPlayer:  2,
			secondPlayer: 4,
			want:         []int{3, 4},
		},
		{
			name:         "Example 2",
			n:            5,
			firstPlayer:  1,
			secondPlayer: 5,
			want:         []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := earliestAndLatest(tt.n, tt.firstPlayer, tt.secondPlayer)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("earliestAndLatest() = %v, want %v", got, tt.want)
			}
		})
	}
}
