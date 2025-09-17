package problems

import (
	"testing"
)

type FoodOp struct {
	op      string
	food    string
	cuisine string
	rating  int
	want    string
}

func TestFoodRatings(t *testing.T) {
	tests := []struct {
		name     string
		foods    []string
		cuisines []string
		ratings  []int
		ops      []FoodOp
	}{
		{
			name:     "Example 1",
			foods:    []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
			cuisines: []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
			ratings:  []int{9, 12, 8, 15, 14, 7},
			ops: []FoodOp{
				{op: "highestRated", cuisine: "korean", want: "kimchi"},
				{op: "highestRated", cuisine: "japanese", want: "ramen"},
				{op: "changeRating", food: "sushi", rating: 16},
				{op: "highestRated", cuisine: "japanese", want: "sushi"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FR := FoodRatingsConstructor(tt.foods, tt.cuisines, tt.ratings)
			for _, op := range tt.ops {
				switch op.op {
				case "changeRating":
					FR.ChangeRating(op.food, op.rating)
				case "highestRated":
					got := FR.HighestRated(op.cuisine)
					if got != op.want {
						t.Errorf("foodRatings() = %v, want %v", got, op.want)
					}
				}
			}
		})
	}
}
