package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"mathskills"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide the file name as an argument.")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Slice to store integer values parsed from the lines
	var linesInt []int
	if lines== nil {
		fmt.Println("Error: Empty File")
		return
	}
	// Convert lines to integers
	for _, line := range lines {
		a, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Conversion error:", err)
			continue
		}
		linesInt = append(linesInt, a)
	}

	// print the mean, median, variance, and standard deviation
	fmt.Println("Average:", mathskills.Mean(linesInt))
	fmt.Println("Median:", mathskills.Median(linesInt))
	fmt.Println("Variance:", int(math.Round(math.Sqrt(float64(mathskills.StanDiv(linesInt))))))
	fmt.Println("Standard Deviation:", mathskills.StanDiv(linesInt))
}
