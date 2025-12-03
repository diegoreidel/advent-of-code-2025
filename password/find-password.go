package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
	"strconv"
)

func main() {

	startingPosition := 50

	lines, err := io.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	firstPuzzle(startingPosition, lines)
	secondPuzzle(startingPosition, lines)
}

func firstPuzzle(startingPosition int, lines []string) {
	position := startingPosition
	password := 0

	for _, line := range lines {
		clicks := readClicks(line)

		clicks = clicks % 100
		position += clicks

		if position < 0 {
			position += 100
		} else if position > 99 {
			position = position % 100
		}

		if position == 0 {
			password++
		}
	}

	fmt.Println("The answer for the first puzzle is: ", password)
}

func secondPuzzle(startingPosition int, lines []string) {
	position := startingPosition
	password := 0

	for _, line := range lines {
		clicks := readClicks(line)

		password += max(clicks, -clicks) / 100
		clicks = clicks % 100
		nextPosition := position + clicks

		if (nextPosition <= 0 || nextPosition > 99) && position != 0 {
			password++
		}

		if nextPosition < 0 {
			position = nextPosition + 100
		} else if nextPosition > 99 {
			position = nextPosition % 100
		} else {
			position = nextPosition
		}
	}

	fmt.Println("The answer for the second puzzle is: ", password)
}

func readClicks(line string) int {
	direction := line[0:1]
	numberOfClicks, err := strconv.Atoi(line[1:len(line)])
	if err != nil {
		panic(err)
	}

	if direction == "R" {
		return numberOfClicks
	}

	return -numberOfClicks
}
