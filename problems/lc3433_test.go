package problems

import (
	"reflect"
	"testing"
)

func TestCountMentions(t *testing.T) {
	tests := []struct {
		name          string
		numberOfUsers int
		events        [][]string
		want          []int
	}{
		{
			name:          "Example 1",
			numberOfUsers: 2,
			events:        [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "71", "HERE"}},
			want:          []int{2, 2},
		},
		{
			name:          "Example 2",
			numberOfUsers: 2,
			events:        [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "12", "ALL"}},
			want:          []int{2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countMentions(tt.numberOfUsers, tt.events)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countMentions() = %v, want %v", got, tt.want)
			}
		})
	}
}
