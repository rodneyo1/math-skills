package mathskills

import (
	"math"
)

func StanDiv(numbers []float64) int {
	if len(numbers) == 0 {
		return 0
	}
	n := Mean(numbers)

	// Calculate the sum of squared differences from the mean
	var squaredDifferences float64
	for _, num := range numbers {
		squaredDifferences += math.Pow(float64(int(num)-n), 2)
	}

	// Calculate the variation
	variation := squaredDifferences / float64(len(numbers))

	stdDev := math.Sqrt(variation)

	return int(math.Round(stdDev))
}
