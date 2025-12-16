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
