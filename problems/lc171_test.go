package problems

import (
	"reflect"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		name        string
		columnTitle string
		want        int
	}{
		{
			name:        "Example 1",
			columnTitle: "A",
			want:        1,
		},
		{
			name:        "Example 2",
			columnTitle: "AA",
			want:        27,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := titleToNumber(tt.columnTitle)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
