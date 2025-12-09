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
	boxA     Box
	boxB     Box
	distance float64
}

type Vertice struct {
	box   Box
	edges []Edge
}

type Box struct {
	x float64
	y float64
	z float64
}

func main() {
	input, err := io.ReadFile("./input/test.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	boxes := buildBoxes(input)
	sortedEdges := calculateDistances(boxes)

	sort.Slice(sortedEdges, func(i, j int) bool {
		return sortedEdges[i].distance < sortedEdges[j].distance
	})

	edges := sortedEdges[:10]
	graph := make(map[string]Vertice)

	for _, edge := range edges {
		boxAString := boxToString(edge.boxA)
		boxBString := boxToString(edge.boxB)
		graph[boxAString] = Vertice{box: edge.boxA}
		graph[boxBString] = Vertice{box: edge.boxB}
	}

	for _, edge := range edges {
		boxAString := boxToString(edge.boxA)
		boxBString := boxToString(edge.boxB)

		if val, exists := graph[boxAString]; exists {
			val.edges = append(val.edges, edge)
			graph[boxAString] = val
		}

		if val, exists := graph[boxBString]; exists {
			val.edges = append(val.edges, edge)
			fmt.Println(val.edges)
			graph[boxBString] = val
		}
	}

	fmt.Println(graph)

	seen := make(map[string]struct{})
	var sizes []int
	for _, vertice := range graph {
		size := countSizeOfCircuit(vertice, graph, seen)
		if size > 0 {
			fmt.Println("Size: ", size)
		}
	}

	fmt.Println("The answer for the first puzzle is: ", sizes)

}

func countSizeOfCircuit(vertice Vertice, graph map[string]Vertice, seen map[string]struct{}) int {
	if _, exists := seen[boxToString(vertice.box)]; exists {
		return 0
	}

	seen[boxToString(vertice.box)] = struct{}{}
	size := 1
	for _, edge := range vertice.edges {
		size += countSizeOfCircuit(graph[boxToString(edge.boxA)], graph, seen)
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
			distances = append(distances, Edge{boxA: boxA, boxB: boxB, distance: distance})
		}
	}

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
