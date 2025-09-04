package problems

import (
	"reflect"
	"testing"
)

func TestFindClosest(t *testing.T) {
	tests := []struct {
		name    string
		x, y, z int
		want    int
	}{
		{
			name: "Example 1",
			x:    2, y: 7, z: 4,
			want: 1,
		},
		{
			name: "Example 2",
			x:    2, y: 5, z: 6,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findClosest(tt.x, tt.y, tt.z)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
