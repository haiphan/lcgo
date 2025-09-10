package problems

import (
	"reflect"
	"testing"
)

func TestMinimumTeachings(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		languages   [][]int
		friendships [][]int
		want        int
	}{
		{
			name:        "Example 1",
			n:           2,
			languages:   [][]int{{1}, {2}, {1, 2}},
			friendships: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:        1,
		},
		{
			name:        "Example 2",
			n:           3,
			languages:   [][]int{{2}, {1, 3}, {1, 2}, {3}},
			friendships: [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}},
			want:        2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumTeachings(tt.n, tt.languages, tt.friendships)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumTeachings() = %v, want %v", got, tt.want)
			}
		})
	}
}
