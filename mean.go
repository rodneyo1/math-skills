package mathskills

import "math"

func Mean(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	var mean float64 = float64(sum / len(numbers))
	return int(math.Round(mean))
}
