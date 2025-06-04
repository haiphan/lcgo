package problems

import (
	"reflect"
	"testing"
)

func TestAnswerString(t *testing.T) {
	tests := []struct {
		name       string
		word       string
		numFriends int
		want       string
	}{
		{
			name:       "Example 1",
			word:       "gggg",
			numFriends: 4,
			want:       "g",
		},
		{
			name:       "Example 2",
			word:       "dbca",
			numFriends: 2,
			want:       "dbc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := answerString(tt.word, tt.numFriends)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerString() = %v, want %v", got, tt.want)
			}
		})
	}
}
