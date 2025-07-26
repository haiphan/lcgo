package problems

import (
	"reflect"
	"testing"
)

func TestMakesquare(t *testing.T) {
	tests := []struct {
		name        string
		matchsticks []int
		want        bool
	}{
		{
			name:        "Example 1",
			matchsticks: []int{1, 1, 2, 2, 2},
			want:        true,
		},
		{
			name:        "Example 2",
			matchsticks: []int{3, 3, 3, 3, 4},
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makesquare(tt.matchsticks)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
