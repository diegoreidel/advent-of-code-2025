package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
	"strconv"
)

type Node struct {
	row int
	col int
}

func main() {
	input, err := io.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	lookupDirections := [][]int{{1, 0}}
	splitDirections := [][]int{{0, -1}, {0, 1}}

	grid := io.BuildGrid(input)
	start := findStartPoint(grid)

	firstPuzzle(grid, start, lookupDirections, splitDirections)
	secondPuzzle(grid, start, lookupDirections, splitDirections)

}

func firstPuzzle(grid [][]string, start Node, lookupDirections [][]int, splitDirections [][]int) {
	splits := make(map[string]struct{})

	numberOfSplits := shootBean(grid, start, lookupDirections, splitDirections, splits)
	fmt.Println("The answer for the first puzzle is: ", numberOfSplits)
}

func secondPuzzle(grid [][]string, start Node, lookupDirections [][]int, splitDirections [][]int) {
	cache := make(map[string]int) // Initializes an empty map

	numberOfSplits := quantunBean(grid, start, lookupDirections, splitDirections, cache)
	fmt.Println("The answer for the second puzzle is: ", numberOfSplits)
}

func quantunBean(grid [][]string, start Node, lookupDirections [][]int, splitDirections [][]int, cache map[string]int) int {
	if start.row == len(grid)-1 {
		return 1
	}

	nodeId := nodeId(start)
	if val, exists := cache[nodeId]; exists {
		return val
	}

	numberOfSplits := 0

	currentRow := start.row
	currentCol := start.col

	for _, direction := range lookupDirections {
		nextNode := Node{currentRow + direction[0], currentCol + direction[1]}

		if nextNode.row < len(grid) && nextNode.col < len(grid[0]) {
			char := grid[nextNode.row][nextNode.col]
			if char == "^" {

				for _, splitDirection := range splitDirections {
					splitNode := Node{nextNode.row + splitDirection[0], nextNode.col + splitDirection[1]}
					if validNode(grid, splitNode) {
						numberOfSplits += quantunBean(grid, splitNode, lookupDirections, splitDirections, cache)
					}
				}
			} else {
				numberOfSplits += quantunBean(grid, nextNode, lookupDirections, splitDirections, cache)
			}
		}
	}

	cache[nodeId] = numberOfSplits
	return numberOfSplits

}

func shootBean(grid [][]string, start Node, lookupDirections [][]int, splitDirections [][]int, splits map[string]struct{}) int {
	if start.row == len(grid)-1 {
		return 0
	}

	numberOfSplits := 0

	currentRow := start.row
	currentCol := start.col

	for _, direction := range lookupDirections {

		didSplit := false
		nextNode := Node{currentRow + direction[0], currentCol + direction[1]}

		if nextNode.row < len(grid) && nextNode.col < len(grid[0]) {
			char := grid[nextNode.row][nextNode.col]
			if char == "^" {

				for _, splitDirection := range splitDirections {
					splitNode := Node{nextNode.row + splitDirection[0], nextNode.col + splitDirection[1]}
					nodeId := nodeId(splitNode)
					_, exists := splits[nodeId]
					if validNode(grid, splitNode) && !exists {
						didSplit = true
						splits[nodeId] = struct{}{}
						numberOfSplits += shootBean(grid, splitNode, lookupDirections, splitDirections, splits)
					}
				}
			} else {
				numberOfSplits += shootBean(grid, nextNode, lookupDirections, splitDirections, splits)
			}
		}

		if didSplit {
			numberOfSplits++
		}
	}

	return numberOfSplits

}

func validNode(grid [][]string, node Node) bool {
	return node.row >= 0 && node.row < len(grid) && node.col >= 0 && node.col < len(grid[0])
}

func findStartPoint(grid [][]string) Node {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == "S" {
				return Node{row, col}
			}
		}
	}

	return Node{-1, -1}
}

func nodeId(node Node) string {
	return string(strconv.Itoa(node.row) + "-" + strconv.Itoa(node.col))
}
