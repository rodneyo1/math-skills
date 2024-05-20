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

	if t%2 == 0 {
		v := numbers[t/2] + numbers[(t/2)-1]
		return int(math.Round(v / 2))
	}
	return int(math.Round(numbers[t/2]))
}
