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

	firstPuzzle(input[0])
	secondPuzzle(input[0])
}

func firstPuzzle(input string) {

	sum := 0

	for _, line := range strings.Split(input, ",") {
		interval := strings.Split(line, "-")

		low, _ := strconv.Atoi(interval[0])
		high, _ := strconv.Atoi(interval[1])

		for i := low; i <= high; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 == 0 {
				firstHalf := s[:len(s)/2]
				secondHalf := s[len(s)/2:]

				if firstHalf == secondHalf {
					fmt.Println(i)
					sum += i
				}
			}
		}
	}

	fmt.Println("The answer for the first puzzle is: ", sum)

}

func secondPuzzle(input string) {

	sum := 0

	invalids := make(map[int]struct{})

	for _, line := range strings.Split(input, ",") {
		interval := strings.Split(line, "-")

		low, _ := strconv.Atoi(interval[0])
		high, _ := strconv.Atoi(interval[1])

		for number := low; number <= high; number++ {
			s := strconv.Itoa(number)
			for size := 1; size < len(s)/2+1; size++ {
				cutpoint := 0
				set := make(map[string]struct{})
				for cutpoint+size <= len(s) {
					value := s[cutpoint : cutpoint+size]
					set[value] = struct{}{}
					cutpoint += size
				}
				if cutpoint == len(s) && len(set) == 1 {
					invalids[number] = struct{}{}
				}
			}
		}
	}

	fmt.Println("the invalids are: ", invalids)

	for item := range invalids {
		sum += item
	}

	fmt.Println("The answer for the first puzzle is: ", sum)

}
