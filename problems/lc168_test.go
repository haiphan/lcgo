package problems

import (
	"reflect"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		name         string
		columnNumber int
		want         string
	}{
		{
			name:         "Example 1",
			columnNumber: 1,
			want:         "A",
		},
		{
			name:         "Example 2",
			columnNumber: 28,
			want:         "AB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertToTitle(tt.columnNumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
