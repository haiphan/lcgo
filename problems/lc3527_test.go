package problems

import (
	"reflect"
	"testing"
)

func TestFindCommonResponse(t *testing.T) {
	tests := []struct {
		name      string
		responses [][]string
		want      string
	}{
		{
			name:      "Example 1",
			responses: [][]string{{"good", "ok", "good", "ok"}, {"ok", "bad", "good", "ok", "ok"}, {"good"}, {"bad"}},
			want:      "good",
		},
		{
			name:      "Example 2",
			responses: [][]string{{"good", "ok", "good"}, {"ok", "bad"}, {"bad", "notsure"}, {"great", "good"}},
			want:      "bad",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCommonResponse(tt.responses)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCommonResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
