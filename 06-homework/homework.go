package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := io.ReadFile("./input/input.txt")
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
	operations := input[len(input)-1]

	largestSize := 0

	for _, line := range input {
		length := len(line)
		if length > largestSize {
			largestSize = length
		}
	}

	numbers := make([]string, largestSize)

	for col := largestSize - 1; col >= 0; col-- {
		for row := 0; row < len(input)-1; row++ {
			if col < len(input[row]) {
				char := input[row][col]
				if char != ' ' {
					numbers[col] += string(char)
				}
			}
		}
	}

	var operation uint8
	var answer int64
	var aggregation int64

	for i, number := range numbers {
		value, _ := strconv.Atoi(number)
		if i < len(operations) && operations[i] != ' ' {
			operation = operations[i]
			if operation == '*' {
				answer = 1
			}
		}

		if len(strings.TrimSpace(number)) > 0 {
			if operation == '+' {
				answer += int64(value)
			} else if operation == '*' {
				answer *= int64(value)
			}
		} else {
			aggregation += answer
			answer = 0
		}
	}

	aggregation += answer
	fmt.Println("The answer for the second puzzle is: ", aggregation)
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
