# math-skills
Math Skills Calculator

This is a command-line program written in Go that calculates the mean, median, variance, and standard deviation of a list of numbers provided in a text file.
Usage

To use the program, follow these steps:
    Add you data to the data.txt file.
    Navigate to the directory containing the main function.
    Run the program with the main.go file name as an argument: go run main.go [file name]

Example:

bash

go run. data.txt

File Format

The input file should contain a list of numbers, each on a separate line. The numbers can be integers or floating-point numbers. Empty lines and whitespace-only lines are ignored.

Example file content (numbers.txt):

23
45
17.5
81

Edge Cases and Error Handling

    Empty File: If the input file is empty, the program will print an error message and exit.
    Invalid Numbers: If the input file contains invalid numbers (e.g., non-numeric characters), the program will print an error message for each invalid line and continue processing the remaining lines.
    Unrepresentable Numbers: If any number in the input file exceeds the maximum representable integer value (math.MaxInt), the program will print an error message and exit.

Inaccuracies with Large Numbers

    When dealing with very large numbers, such as those that exceed the range of the int or float64 data types, the program may encounter inaccuracies due to the limitations of finite precision representation in computers. These inaccuracies can arise from round-off errors, loss of significance, or overflow/underflow conditions. To mitigate inaccuracies, consider using arbitrary-precision arithmetic libraries or carefully scaling and normalizing input data.