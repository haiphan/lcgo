package problems

import (
	"reflect"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		limit int
		want  int64
	}{
		{
			name:  "Example 1",
			n:     5,
			limit: 2,
			want:  3,
		},
		{
			name:  "Example 2",
			n:     3,
			limit: 3,
			want:  10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distributeCandies(tt.n, tt.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
