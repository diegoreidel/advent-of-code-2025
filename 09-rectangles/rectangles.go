package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	input, err := io.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	points := buildPoints(input)

	firstPuzzle(points)
	secondPuzzle(points)
}

func secondPuzzle(points []Point) {
	maxX := 0
	maxY := 0

	for _, point := range points {
		if maxX < point.x {
			maxX = point.x
		}
		if maxY < point.y {
			maxY = point.y
		}
	}

	var grid [][]string
	for i := 0; i <= maxX; i++ {
		line := make([]string, maxY+1)
		for i, _ := range line {
			line[i] = "."
		}
		grid = append(grid, line)
	}

	for _, point := range points {
		grid[point.x][point.y] = "#"
	}

	for i := 0; i < len(points)-1; i++ {
		buildLimits(points[i], points[i+1], grid)
	}
	buildLimits(points[0], points[len(points)-1], grid)
	fillGreenTiles(grid)
	writeGrid(grid)
}

func fillGreenTiles(grid [][]string) {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for j < len(grid[i]) && grid[i][j] != "X" && grid[i][j] != "#" {
				j++
			}

			j++

			for j < len(grid[i]) && grid[i][j] == "." {
				grid[i][j] = "X"
				j++
			}
		}
	}
}

func buildLimits(pointA Point, pointB Point, grid [][]string) {
	if pointA.x == pointB.x {
		aY := pointA.y
		bY := pointB.y
		for aY-1 > bY {
			grid[pointA.x][aY-1] = "X"
			aY--
		}

		for bY-1 > aY {
			grid[pointB.x][bY-1] = "X"
			bY--
		}
	} else {
		aX := pointA.x
		bX := pointB.x
		for aX-1 > bX {
			grid[aX-1][pointA.y] = "X"
			aX--
		}

		for bX-1 > aX {
			grid[bX-1][pointB.y] = "X"
			bX--
		}
	}
}

func writeGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func firstPuzzle(points []Point) {
	maximum := float64(0)
	for i, _ := range points {
		for j := i + 1; j < len(points); j++ {
			xDistance := math.Abs(float64(points[i].x - points[j].x + 1))
			yDistance := math.Abs(float64(points[i].y - points[j].y + 1))

			area := xDistance * yDistance
			if area > maximum {
				maximum = area
			}
		}
	}

	fmt.Println("The answer for the first puzzle is: ", strconv.FormatFloat(maximum, 'f', -1, 64))
}

func buildPoints(input []string) []Point {

	var points []Point

	for _, line := range input {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[1])
		y, _ := strconv.Atoi(coords[0])
		points = append(points, Point{x, y})
	}

	return points
}
