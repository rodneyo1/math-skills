package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"

	"mathskills"
)

func main() {

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error Opening File:", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if err != nil {
			break
		}
	}

	// Slice to store integer values parsed from the lines
	var linesInt []int

	// Convert lines to integers
	for i := 0; i < len(lines); i++ {
		a, err := strconv.Atoi(lines[i])
		linesInt = append(linesInt, a)
		if err != nil {
			fmt.Println("Conversion Error:", err)
		}
	}

	fmt.Println("Average:", mathskills.Mean(linesInt))
	fmt.Println("Median:", mathskills.Median(linesInt))
	fmt.Println("Variance:", int(math.Sqrt(mathskills.StanDiv(linesInt))))
	fmt.Println("Standard Deviation:", mathskills.StanDiv(linesInt))
}
