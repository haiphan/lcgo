package problems

import (
	"reflect"
	"testing"
)

func TestMinOperationsQueries(t *testing.T) {
	tests := []struct {
		name    string
		queries [][]int
		want    int64
	}{
		{
			name:    "Example 1",
			queries: [][]int{{1, 2}, {2, 4}},
			want:    3,
		},
		{
			name:    "Example 2",
			queries: [][]int{{2, 6}},
			want:    4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperationsQueries(tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperationsQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
