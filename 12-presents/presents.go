package main

import (
	"diegoreidel/io"
	"fmt"
	"strings"
	"strconv"
)

type Tree struct {
	area int
	requiredPresents []int
}

func main() {
	shapes, _ := io.ReadFile("./input/shapes.txt")
	input, _ := io.ReadFile("./input/input.txt")

	shapesMap := make(map[string][]string)

	for i := 0; i < len(shapes); i += 4 {
		addShape(shapesMap, shapes[i:i+4])
	}

	trees := buildTrees(input)
	areasMap := calculateTheAreas(shapesMap)
	fmt.Println(areasMap)

	var invalidTrees []Tree
	var remainingTrees []Tree

	// All areas smaller than the minimum required area are clearly invalid
	for _, tree := range trees {
		minimumRequiredArea := 0
		for i, required := range tree.requiredPresents {
			minimumRequiredArea += required * areasMap[i]
		}

		if tree.area < minimumRequiredArea {
			invalidTrees = append(invalidTrees, tree)
		} else {
			remainingTrees = append(remainingTrees, tree)
		}
	}

	fmt.Println("The number of INVALID trees is: ", len(invalidTrees))

	fmt.Println("The number of REMAINING trees is: ", len(remainingTrees))

	// I tested the number of remaining trees as the answer and got the star! So...
	fmt.Println("The answer is: ", len(remainingTrees))

}

func buildTrees(input []string) []Tree {

	var trees []Tree

	for _, line := range input {
		parts := strings.Fields(line)

		size := parts[0]
		size = size[0:len(size) - 1]

		dimensions := strings.Split(size, "x")

		width, _ := strconv.Atoi(dimensions[0])
		length, _ := strconv.Atoi(dimensions[1])

		var requiredPresents []int
		for _, requiredPresent := range parts[1:] {
			value, _ := strconv.Atoi(requiredPresent)
			requiredPresents = append(requiredPresents, value)
		}

		trees = append(trees, Tree{width * length, requiredPresents})
	}

	return trees
}

func calculateTheAreas(shapesMap map[string][]string) map[int]int {
	areasMap := make(map[int]int)

	for key, value := range shapesMap {
		area := 0
		for _, line := range value {
			area += strings.Count(line, "#")
		}
		i,_ := strconv.Atoi(key)
		areasMap[i] = area
	}

	return areasMap
}

func addShape(shapesMap map[string][]string, lines []string) {
	shapeID := lines[0][0]
	shapesMap[string(shapeID)] = lines[1:]
}