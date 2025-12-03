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

	firstPuzzle(input)
	secondPuzzle(input)
}

// My second solution can also solve the first puzzle, but because I had initially solved it using this first method, I decided to keep it.
func firstPuzzle(input []string) {

	maximum := 0

	for _, line := range input {
		characters := strings.Split(line, "")
		largestFound := 0
		for pos, number := range characters {
			current, _ := strconv.Atoi(number)
			restOfInput := line[pos+1:]
			positionOfNextLargest := findPositionOfLargest(restOfInput)

			if positionOfNextLargest != -1 {
				nextLargest := int(restOfInput[positionOfNextLargest] - '0')

				calculation := current*10 + nextLargest

				if calculation > largestFound {
					largestFound = calculation
				}
			}
		}

		maximum += largestFound
	}

	fmt.Println("The answer for the first puzzle is: ", maximum)
}

func secondPuzzle(input []string) {

	maximum := 0

	for _, line := range input {
		largest := findLargestNumberPossible(line, 11)

		numeric, _ := strconv.Atoi(largest)
		maximum += numeric
	}

	fmt.Println("The answer for the second puzzle is: ", maximum)
}

func findLargestNumberPossible(line string, removeLast int) string {
	if removeLast < 0 {
		return ""
	}

	largestPosition := findPositionOfLargest(line[:len(line)-removeLast])
	return string(line[largestPosition]) + findLargestNumberPossible(line[largestPosition+1:], removeLast-1)
}

func findPositionOfLargest(line string) int {
	characters := strings.Split(line, "")
	largest := -1
	for pos, _ := range characters {
		if largest == -1 || line[pos] > line[largest] {
			largest = pos
		}
	}

	return largest
}
