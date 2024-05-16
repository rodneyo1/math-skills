package mathskills

import "math"

func Mean(numbers []float64) int {
	if len(numbers) == 0 {
		return 0
	}
	var sum float64
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	var mean float64 = sum / float64(len(numbers))

	return int(math.Round(mean))
}
