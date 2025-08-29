package problems

import (
	"reflect"
	"testing"
)

func TestFlowerGame(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		want int64
	}{
		{
			name: "Example 1",
			n:    2,
			m:    3,
			want: 3,
		},
		{
			name: "Example 2",
			n:    1,
			m:    1,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flowerGame(tt.n, tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flowerGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
