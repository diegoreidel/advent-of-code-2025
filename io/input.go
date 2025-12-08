package io

import (
	"bufio"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func BuildGrid(input []string) [][]string {
	rows := len(input)
	cols := len(strings.Fields(input[0]))

	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, cols)
	}

	for row, line := range input {
		data := strings.Split(line, "")
		grid[row] = data
	}

	return grid
}
