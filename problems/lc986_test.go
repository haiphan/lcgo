package problems

import (
	"reflect"
	"testing"
)

func TestIntervalIntersection(t *testing.T) {
	tests := []struct {
		name       string
		firstList  [][]int
		secondList [][]int
		want       [][]int
	}{
		{
			name:       "Example 1",
			firstList:  [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			secondList: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			want:       [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			name:       "Example 2",
			firstList:  [][]int{{1, 3}, {5, 9}},
			secondList: [][]int{},
			want:       [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intervalIntersection(tt.firstList, tt.secondList)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
