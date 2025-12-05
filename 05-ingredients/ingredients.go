package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Range struct {
	low  int
	high int
}

func main() {
	input, err := io.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var ranges []Range
	var ingredients []int
	readingRanges := true

	for _, line := range input {
		if line == "" {
			readingRanges = false
		} else if readingRanges {
			boundaries := strings.Split(line, "-")
			low, _ := strconv.Atoi(boundaries[0])
			high, _ := strconv.Atoi(boundaries[1])

			ranges = append(ranges, Range{low, high})
		} else {
			i, _ := strconv.Atoi(line)
			ingredients = append(ingredients, i)
		}
	}

	firstPuzzle(ranges, ingredients)
	secondPuzzle(ranges)
}

func firstPuzzle(ranges []Range, ingredients []int) {
	numberOfValidIngredients := 0
	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if r.low <= ingredient && r.high >= ingredient {
				numberOfValidIngredients++
				break
			}
		}
	}
	fmt.Println("The answer for the first puzzle is: ", numberOfValidIngredients)
}

func secondPuzzle(ranges []Range) {
	sizeBeforeCleanup := len(ranges)
	cleanedRanges := cleanOverlaps(ranges)

	for sizeBeforeCleanup > len(cleanedRanges) {
		sizeBeforeCleanup = len(cleanedRanges)
		cleanedRanges = cleanOverlaps(cleanedRanges)
	}

	answer := 0
	for _, r := range cleanedRanges {
		answer += r.high - r.low + 1
	}

	fmt.Println("The answer for the second puzzle is: ", answer)
}

func cleanOverlaps(ranges []Range) []Range {
	var cleaned []Range

	for _, r := range ranges {
		wasMerged := false
		for pos, c := range cleaned {
			if overlap(r, c) {
				merged := merge(r, c)
				cleaned[pos] = merged
				wasMerged = true
				break
			}
		}
		if !wasMerged {
			cleaned = append(cleaned, r)
		}
	}

	return cleaned
}

func overlap(range1 Range, range2 Range) bool {
	return range1.low <= range2.high && range2.low <= range1.high
}

func merge(range1 Range, range2 Range) Range {
	return Range{min(range1.low, range2.low), max(range1.high, range2.high)}
}
