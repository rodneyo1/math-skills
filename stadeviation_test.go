package mathskills

import (
  "testing"
)

func TestStanDiv(t *testing.T) {
  // Define some test cases
  tests := []struct {
    name string
    numbers []int
    expected int
  }{
    {name: "single element", numbers: []int{5}, expected: 0},  // Standard deviation of single element is 0
    {name: "positive numbers", numbers: []int{1, 2, 3}, expected: 1},
    {name: "negative numbers", numbers: []int{-1, -2, -3}, expected: 1},
    {name: "mixed numbers", numbers: []int{1, -2, 3}, expected: 2},
    {name: "large numbers", numbers: []int{100, 200, 300}, expected: 82},
  }

  for _, tc := range tests {
    t.Run(tc.name, func(t *testing.T) {
      actual := StanDiv(tc.numbers)
      if actual != tc.expected {
        t.Errorf("Expected standard deviation of %v to be %d, got %d", tc.numbers, tc.expected, actual)
      }
    })
  }
}