package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := io.ReadFile("./input/test.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	grid := buildGrid(input)

	firstPuzzle(grid)
	secondPuzzle(input)
}

func firstPuzzle(grid [][]string) {
	operations := grid[len(grid)-1]
	cols := len(operations)

	var aggregation int64

	for col := 0; col < cols; col++ {

		var answer int64
		answer = 0
		operation := operations[col]
		for row := 0; row < len(grid)-1; row++ {
			value, _ := strconv.Atoi(grid[row][col])
			if row == 0 && operation == "*" {
				answer = 1
			}

			if operation == "+" {
				answer += int64(value)
			} else if operation == "*" {
				answer *= int64(value)
			}
		}

		aggregation += answer
	}

	fmt.Println("The answer for the first puzzle is: ", aggregation)
}

func secondPuzzle(input []string) {

	numberOfLines := 0
	var sizes []int
	largestSize := 0

	for _, line := range input {
		numberOfLines++
		length := len(line)
		sizes = append(sizes, length)
		if length > largestSize {
			largestSize = length
		}
	}

	numbers := make([]string, largestSize)
	operations := ""

	for col := largestSize - 1; col >= 0; col-- {
		for row := 0; row < numberOfLines; row++ {
			if col < len(input[row]) {
				char := input[row][col]
				if char == '+' || char == '*' {
					operations = string(char) + operations
				} else if char != ' ' {
					numbers[col] += string(char)
				}
			}
		}
	}

	var noSpaces []string
	for _, number := range numbers {
		if number != "" {
			noSpaces = append(noSpaces, number)
		}
	}

	fmt.Println("The numbers are: ", noSpaces)
	fmt.Println("Operations are: ", operations)

	values := len(input) - 1

	fmt.Println("The the number of values per operation is : ", values)

	var operation uint8
	var answer int64
	var aggreation int64

	for i, number := range noSpaces {
		value, _ := strconv.Atoi(number)

		if i%values == 0 {
			fmt.Println("The answer for group ", i/values, "is", answer)
			aggreation += answer
			operation = operations[i/values]

			if operation == '*' {
				answer = 1
			} else {
				answer = 0
			}
		}

		if operation == '+' {
			answer += int64(value)
		} else if operation == '*' {
			answer *= int64(value)
		}
	}

	fmt.Println("The answer for the last group is", answer)

	aggreation += answer

	fmt.Println("The answer for the second puzzle is: ", aggreation)

}

func buildGrid(input []string) [][]string {
	rows := len(input)
	cols := len(strings.Fields(input[0]))

	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, cols)
	}

	for row, line := range input {
		data := strings.Fields(line)
		grid[row] = data
	}

	return grid
}
