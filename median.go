package mathskills

import (
	"math"
	"sort"
)

func Median(numbers []int) int {
	sort.Ints(numbers) // Sort the numbers in ascending order
	t := len(numbers)

	// Check if the number of elements is even or odd
	if t%2 == 0 {
		v := float64(numbers[t/2] + numbers[(t/2)-1])
		median := v / 2
		return int(math.Round(median))
	} else {
		return numbers[t/2]
	}
}
