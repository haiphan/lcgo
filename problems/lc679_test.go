package problems

import (
	"reflect"
	"testing"
)

func TestJudgePoint24(t *testing.T) {
	tests := []struct {
		name  string
		cards []int
		want  bool
	}{
		{
			name:  "Example 1",
			cards: []int{4, 1, 8, 7},
			want:  true,
		},
		{
			name:  "Example 2",
			cards: []int{1, 2, 1, 2},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := judgePoint24(tt.cards)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("judgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}
