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
			stringNumber := strconv.Itoa(i)
			if len(stringNumber)%2 == 0 {
				firstHalf := stringNumber[:len(stringNumber)/2]
				secondHalf := stringNumber[len(stringNumber)/2:]

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

	answer := 0

	for _, line := range strings.Split(input, ",") {
		interval := strings.Split(line, "-")

		low, _ := strconv.Atoi(interval[0])
		high, _ := strconv.Atoi(interval[1])

		for number := low; number <= high; number++ {
			stringNumber := strconv.Itoa(number)
			for size := 1; size < len(stringNumber)/2+1; size++ {
				cutPoint := 0
				setOfPossibleSubStrings := make(map[string]struct{})

				for cutPoint+size <= len(stringNumber) {
					value := stringNumber[cutPoint : cutPoint+size]
					setOfPossibleSubStrings[value] = struct{}{}
					cutPoint += size
				}

				if cutPoint == len(stringNumber) && len(setOfPossibleSubStrings) == 1 {
					answer += number
					break
				}
			}
		}
	}

	fmt.Println("The answer for the first puzzle is: ", answer)

}
