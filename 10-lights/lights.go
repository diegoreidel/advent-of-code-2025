package main

import (
	"diegoreidel/io"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Machine struct {
	goal    string
	buttons []Button
	joltage bool
}

type Button struct {
	changes []int
}

func main() {
	input, err := io.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var machines []Machine
	for _, line := range input {
		machine := buildMachine(line, false)
		machines = append(machines, machine)
	}

	firstPuzzle(machines)
}

func firstPuzzle(machines []Machine) {
	answer := 0
	for i, machine := range machines {
		seen := make(map[string]struct{})
		minimal := dp(machine.goal, initialState(machine), machine.buttons, copyMap(seen), []Button{}, math.MaxInt/2, make(map[string]int))
		fmt.Println("For machine ", i+1, " the minimal cost is ", minimal)
		answer += minimal
	}

	fmt.Println("The answer for the first puzzle is: ", answer)
}

func initialState(machine Machine) string {
	initial := ""

	for i := 0; i < len(machine.goal); i++ {
		if machine.joltage {
			initial += "0"
		} else {
			initial += "."
		}
	}

	return initial
}

func dp(goal string, state string, buttons []Button, seen map[string]struct{}, pressed []Button, bestSoFar int, minimals map[string]int) int {

	if state == goal {
		//fmt.Println(pressed)
		return 0
	}

	if val, exists := minimals[state]; exists {
		return val
	}

	if _, exists := seen[state]; exists {
		return math.MaxInt / 2
	}

	if len(pressed) >= bestSoFar {
		return math.MaxInt / 2
	}

	seen[state] = struct{}{}

	answer := math.MaxInt / 2

	for _, button := range buttons {
		newLights := changeLights(state, button.changes)

		copied := copyList(pressed)
		copied = append(copied, button)
		cost := dp(goal, newLights, buttons, copyMap(seen), copied, answer, minimals) + 1
		if cost < answer {
			answer = cost
		}
	}

	minimals[state] = answer

	return answer
}

func copyList(pressed []Button) []Button {
	copied := make([]Button, len(pressed))
	copy(copied, pressed)
	return copied
}

func changeLights(lights string, changes []int) string {

	runes := []rune(lights)
	for _, change := range changes {
		if runes[change] == '.' {
			runes[change] = '#'
		} else {
			runes[change] = '.'
		}
	}

	changed := string(runes)
	return changed
}

func buildMachine(line string, useJoltage bool) Machine {
	parts := strings.Split(line, " ")
	lights := parts[0]
	//joltage := parts[len(parts)-1]

	var buttons []Button

	for _, b := range parts[1 : len(parts)-1] {

		button := Button{changes: []int{}}

		b = b[1 : len(b)-1]
		lightPositions := strings.Split(b, ",")

		for _, pos := range lightPositions {
			value, _ := strconv.Atoi(pos)
			button.changes = append(button.changes, value)
		}

		buttons = append(buttons, button)
	}

	return Machine{lights[1 : len(lights)-1], buttons, useJoltage}
}

func copyMap(original map[string]struct{}) map[string]struct{} {
	copiedMap := make(map[string]struct{}, len(original))
	for key, value := range original {
		copiedMap[key] = value
	}

	return copiedMap
}
