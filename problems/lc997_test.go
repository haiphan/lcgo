package problems

import (
	"reflect"
	"testing"
)

func TestFindJudge(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		trust [][]int
		want  int
	}{
		{
			name:  "Example 1",
			n:     2,
			trust: [][]int{{1, 2}},
			want:  2,
		},
		{
			name:  "Example 2",
			n:     3,
			trust: [][]int{{1, 3}, {2, 3}},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findJudge(tt.n, tt.trust)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
