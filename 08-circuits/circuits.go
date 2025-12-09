package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	source   Box
	target   Box
	distance float64
}

type Box struct {
	x float64
	y float64
	z float64
}

func main() {
	input, err := io.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	boxes := buildBoxes(input)
	sortedEdges := calculateDistances(boxes)

	firstPuzzle(sortedEdges[:1000], 3)
	secondPuzzle(sortedEdges, len(boxes))

}

func firstPuzzle(edges []Edge, k int) {
	seen := make(map[string]struct{})

	sizes := buildGraphAndCountSizes(edges, seen)
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	answer := 1
	for _, size := range sizes[:k] {
		answer *= size
	}

	fmt.Println("The answer for the first puzzle is: ", answer)
}

func secondPuzzle(edges []Edge, numberOfBoxes int) {
	pivot := numberOfBoxes / 2

	seen := make(map[string]struct{})
	size := buildGraphAndCountSizes(edges[:pivot], seen)

	for len(size) == 1 {
		seen = make(map[string]struct{})
		pivot = pivot / 2
		size = buildGraphAndCountSizes(edges[:pivot], seen)
	}
	for len(seen) < numberOfBoxes {
		pivot++
		size = buildGraphAndCountSizes(edges[:pivot], seen)
	}

	edge := edges[pivot-1]

	fmt.Println("The answer for the second puzzle is: ", int(edge.source.x)*int(edge.target.x))
}

func buildGraphAndCountSizes(edges []Edge, seen map[string]struct{}) []int {

	graph := make(map[string][]string)

	for _, edge := range edges {

		sourceId := boxToString(edge.source)
		targetId := boxToString(edge.target)

		if _, exists := graph[sourceId]; !exists {
			graph[sourceId] = []string{}
		}

		if _, exists := graph[targetId]; !exists {
			graph[targetId] = []string{}
		}

		graph[sourceId] = append(graph[sourceId], targetId)
		graph[targetId] = append(graph[targetId], sourceId)

	}

	var sizes []int
	for key, _ := range graph {
		if _, exists := seen[key]; !exists {
			size := countSizeOfCircuit(key, graph, seen)
			sizes = append(sizes, size)
		}
	}

	return sizes
}

func countSizeOfCircuit(key string, graph map[string][]string, seen map[string]struct{}) int {
	if _, exists := seen[key]; exists {
		return 0
	}

	seen[key] = struct{}{}
	size := 1
	for _, edge := range graph[key] {
		size += countSizeOfCircuit(edge, graph, seen)
	}

	return size
}

func calculateDistances(boxes []Box) []Edge {

	var distances []Edge

	for i := 0; i < len(boxes)-1; i++ {
		boxA := boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			boxB := boxes[j]
			distance := math.Sqrt(math.Pow((boxA.x-boxB.x), 2) + math.Pow((boxA.y-boxB.y), 2) + math.Pow((boxA.z-boxB.z), 2))
			distances = append(distances, Edge{source: boxA, target: boxB, distance: distance})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	return distances
}

func buildBoxes(input []string) []Box {

	var boxes []Box

	for _, line := range input {
		coordinates := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(coordinates[0], 64)
		y, _ := strconv.ParseFloat(coordinates[1], 64)
		z, _ := strconv.ParseFloat(coordinates[2], 64)
		boxes = append(boxes, Box{x: x, y: y, z: z})
	}

	return boxes

}

func boxToString(box Box) string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f)", box.x, box.y, box.z)
}
