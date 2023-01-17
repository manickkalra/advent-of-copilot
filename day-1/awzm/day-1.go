package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// sum the numbers in a list
func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// return the largest number in a list
func largest(numbers []int) int {
	largest := 0
	for _, number := range numbers {
		if number > largest {
			largest = number
		}
	}
	return largest
}

// parseEmptyLines uses strings.Split to parse a string into a list of strings
// by the empty lines
func parseEmptyLines(input string) []string {
	return strings.Split(input, "\n\n")
}

// parseNumbers uses strings.Split to parse a string into a int list
// split by each new line
func parseNumbers(input string) []int {
	// parse input into a list of strings
	lines := strings.Split(input, "\n")

	numbers := make([]int, len(lines))
	for i, line := range lines {

		// cast line as int
		n, err := strconv.Atoi(line)
		// handle error
		if err != nil {
			fmt.Errorf("error parsing %s as int", line)
			os.Exit(1)
		}

		numbers[i] = n
	}

	return numbers
}

// parseInput uses parseEmptyLines and parseNumbers to parse a string
// into a list of int lists
func parseInput(input string) [][]int {
	// parse input into a list of strings
	emptyLines := parseEmptyLines(input)

	// parse each empty line into a list of numbers
	parsed := make([][]int, len(emptyLines))
	for i, emptyLine := range emptyLines {
		parsed[i] = parseNumbers(emptyLine)

	}

	return parsed
}

func findLargest(input string) int {

	parsed := parseInput(input)

	largest := 0
	for _, p := range parsed {
		sum := sum(p)

		// set largest to the greater between largest and sum
		if sum > largest {
			largest = sum
		}
	}

	return largest
}

// read file and return the contents as a string
func readFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		panic(err)
	}

	return string(bs)
}

// main func calls "findLargest" and prints the result
func main() {
	// read input from input.txt as a string
	input := readFile("input.txt")
	result := findLargest(input)
	println(result)
}
