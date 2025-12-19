package problems

import (
	"reflect"
	"testing"
)

func TestFindAllPeople(t *testing.T) {
	tests := []struct {
		name           string
		n, firstPerson int
		meetings       [][]int
		want           []int
	}{
		{
			name:        "Example 1",
			n:           6,
			meetings:    [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}},
			firstPerson: 1,
			want:        []int{0, 1, 2, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAllPeople(tt.n, tt.meetings, tt.firstPerson)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
