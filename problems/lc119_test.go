package problems

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{
			name:     "Example 1",
			rowIndex: 3,
			want:     []int{1, 3, 3, 1},
		},
		{
			name:     "Example 2",
			rowIndex: 0,
			want:     []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRow(tt.rowIndex)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
