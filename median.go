package mathskills

import (
	"math"
	"sort"
)

func Median(numbers []float64) int {
	if len(numbers) == 0 {
		return 0
	}
	sort.Float64s(numbers) // Sort the numbers in ascending order
	t := len(numbers)

	// Check if the number of elements is even or odd
	if t%2 == 0 {
		v := float64(numbers[t/2] + numbers[(t/2)-1])
		median := v / 2
		return int(math.Round(median))
	} else {
		return int(math.Round(numbers[t/2]))
	}
}
