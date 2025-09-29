package problems

import (
	"reflect"
	"testing"
)

func TestMinScoreTriangulation(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "Example 1",
			values: []int{1, 2, 3},
			want:   6,
		},
		{
			name:   "Example 2",
			values: []int{3, 7, 4, 5},
			want:   144,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minScoreTriangulation(tt.values)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minScoreTriangulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
