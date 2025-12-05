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

	var ranges []string
	var ingredients []string
	readingRanges := true

	for _, line := range input {
		if line == "" {
			readingRanges = false
		} else if readingRanges {
			ranges = append(ranges, line)
		} else {
			ingredients = append(ingredients, line)
		}
	}

	firstPuzzle(ranges, ingredients)
}

func firstPuzzle(ranges []string, ingredients []string) {

	validIngredients := make(map[string]bool)

	for _, r := range ranges {
		boundaries := strings.Split(r, "-")
		low, _ := strconv.Atoi(boundaries[0])
		high, _ := strconv.Atoi(boundaries[1])

		for i := low; i <= high; i++ {
			validIngredients[strconv.Itoa(i)] = true
		}
	}

	numberOfValidIngredients := countValidIngredients(validIngredients, ingredients)
	fmt.Println("The answer for the first puzzle is: ", numberOfValidIngredients)
}

func countValidIngredients(valid map[string]bool, ingredients []string) int {

	numberOfValidIngredients := 0
	for _, i := range ingredients {
		if valid[i] {
			numberOfValidIngredients++
		}
	}

	return numberOfValidIngredients
}
