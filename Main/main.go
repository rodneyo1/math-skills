package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"mathskills"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: [Program Name] [File name]")
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
	var linesInt []float64
	if lines == nil {
		fmt.Println("Error: Empty File")
		return
	}
	// Convert lines to integers
	for indx, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		a, err := strconv.ParseFloat(trimmedLine, 64)
		if err != nil {
			fmt.Println("Conversion error:", err)
			continue
		}
		if a > float64(math.MaxInt) {
			//You can encounter unrepresentable numbers even if they are within the range of math.MaxFloat64
			fmt.Printf("Unrepresentable number encountered at line %d\n", indx+1)
			return
		}
		linesInt = append(linesInt, a)
	}

	// print the mean, median, variance, and standard deviation
	fmt.Println("Average:", mathskills.Mean(linesInt))
	fmt.Println("Median:", mathskills.Median(linesInt))
	fmt.Println("Variance:", int(math.Round(math.Sqrt(float64(mathskills.StanDiv(linesInt))))))
	fmt.Println("Standard Deviation:", mathskills.StanDiv(linesInt))
}
