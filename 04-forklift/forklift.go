package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
)

func main() {
	input, err := io.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	firstPuzzle(input, directions)
	secondPuzzle(input, directions)
}

func firstPuzzle(input []string, directions [][]int) {
	answer := 0
	grid := buildGrid(input)
	afterCleanUp := deepCopy(grid)

	answer += countAndRemoveRolls(grid, directions, afterCleanUp)
	fmt.Println("The answer for the first puzzle is: ", answer)
}

func secondPuzzle(input []string, directions [][]int) {
	previousAnswer := 0
	answer := 0
	afterCleanUp := buildGrid(input)

	for previousAnswer == 0 || previousAnswer != answer {
		currentGrid := afterCleanUp
		afterCleanUp = deepCopy(currentGrid)

		previousAnswer = answer
		answer += countAndRemoveRolls(currentGrid, directions, afterCleanUp)
	}

	fmt.Println("The answer for the second puzzle is: ", answer)
}

func countAndRemoveRolls(currentGrid [][]int32, directions [][]int, cleaned [][]int32) int {
	answer := 0

	for row := 0; row < len(currentGrid); row++ {
		for col := 0; col < len(currentGrid[row]); col++ {
			current := currentGrid[row][col]
			if current == '@' {
				adjacent := countAdjacent(currentGrid, row, col, directions)
				if adjacent < 4 {
					answer++
					cleaned[row][col] = '.'
				}
			}
		}
	}
	return answer
}

func countAdjacent(grid [][]int32, row int, col int, directions [][]int) int {
	adjacentRolls := 0

	for direction := range directions {
		adjacentRow := row + directions[direction][0]
		adjacentCol := col + directions[direction][1]
		if adjacentRow >= 0 && adjacentRow < len(grid) && adjacentCol >= 0 && adjacentCol < len(grid[0]) && grid[adjacentRow][adjacentCol] == '@' {
			adjacentRolls++
		}
	}

	return adjacentRolls
}

func buildGrid(input []string) [][]int32 {
	rows := len(input)
	cols := len(input[0])

	grid := make([][]int32, rows)
	for i := range grid {
		grid[i] = make([]int32, cols)
	}

	for row, line := range input {
		thisRow := grid[row]
		for col, char := range line {
			thisRow[col] = char
		}
	}

	return grid
}

func deepCopy(grid [][]int32) [][]int32 {
	destinationMatrix := make([][]int32, len(grid))

	for i, row := range grid {
		destinationMatrix[i] = make([]int32, len(row))
		copy(destinationMatrix[i], row)
	}

	return destinationMatrix
}
