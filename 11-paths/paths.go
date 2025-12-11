package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := io.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	graph := buildGraph(input)
	paths := search(graph, "you", make(map[string]struct{}), []string{})
	fmt.Println(paths)

	paths = search(graph, "svr", make(map[string]struct{}), []string{"fft", "dac"})
	fmt.Println(paths)

}

func search(graph map[string][]string, current string, seen map[string]struct{}, requiredNodes []string) int  {

	if current == "out" {
		if meetsRequirements(requiredNodes, seen) {
			return 1
		}
	}

	if _, exists := seen[current]; exists {
		return 0
	}

	seen[current] = struct{}{}
	seenClone := copyMap(seen)
	answer := 0
	for _, next := range graph[current] {
		if _, exists := seen[current]; exists {
			answer += search(graph, next, seenClone, requiredNodes)
		}
	}

	return answer
}

func meetsRequirements(requirements []string, seen map[string]struct{}) bool {
	var answer = true
	for _, node := range requirements {
		if _, ok := seen[node]; !ok {
			answer = false
		}
	}
	return answer
}

func buildGraph(input []string) map[string][]string {
	var graph = make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, " ")
		key := parts[0][:len(parts[0])-1]

		graph[key] = parts[1:]

	}

	return graph
}

func copyMap(original map[string]struct{}) map[string]struct{} {
	copiedMap := make(map[string]struct{}, len(original))
	for key, value := range original {
		copiedMap[key] = value
	}

	return copiedMap
}