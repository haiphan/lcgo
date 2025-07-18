package problems

import (
	"reflect"
	"testing"
)

func TesAsteroidCollision(t *testing.T) {
	tests := []struct {
		name      string
		asteroids []int
		want      []int
	}{
		{
			name:      "Example 1",
			asteroids: []int{5, 10, -5},
			want:      []int{5, 10},
		},
		{
			name:      "Example 2",
			asteroids: []int{8, -8},
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := asteroidCollision(tt.asteroids)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
