package mathskills_test

import (
	"mathskills"
	"testing"
)

func TestMean(t *testing.T) {
	// Define some test cases
	tests := []struct {
		name     string
		numbers  []float64
		expected int
	}{

		{name: "single element", numbers: []float64{5}, expected: 5},
		{name: "positive numbers", numbers: []float64{1, 2, 3}, expected: 2},
		{name: "negative numbers", numbers: []float64{-1, -2, -3}, expected: -2},
		{name: "mixed numbers", numbers: []float64{1, -2, 3}, expected: 1},
		{name: "large numbers", numbers: []float64{100, 200, 300}, expected: 200},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := mathskills.Mean(tc.numbers)
			if actual != tc.expected {
				t.Errorf("Expected mean of %v to be %d, got %d", tc.numbers, tc.expected, actual)
			}
		})
	}
}
