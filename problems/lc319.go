package problems

import "math"

func bulbSwitch(n int) int {
	// Only squares have an odd number of divisors, which means they will remain on after all toggles.
	return int(math.Sqrt(float64(n)))
}
