package problems

import (
	"reflect"
	"testing"
)

func TestCountCollisions(t *testing.T) {
	tests := []struct {
		name       string
		directions string
		want       int
	}{
		{
			name:       "Example 1",
			directions: "RLRSLL",
			want:       5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countCollisions(tt.directions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countCollisions() = %v, want %v", got, tt.want)
			}
		})
	}
}
