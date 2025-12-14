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

	var lights []Machine
	for _, line := range input {
		machine := buildMachine(line, false)
		lights = append(lights, machine)
	}

	firstPuzzle(lights)
}

func firstPuzzle(machines []Machine) {
	answer := 0
	for _, machine := range machines {
		seen := make(map[string]struct{})
		minimal := dp(machine, initialState(machine), machine.buttons, copyMap(seen), []Button{}, math.MaxInt/2, make(map[string]int))
		answer += minimal
	}

	fmt.Println("The answer for the first puzzle is: ", answer)
}

func initialState(machine Machine) string {
	initial := ""

	size := len( machine.goal)
	if machine.joltage {
		size = len(strings.Split(machine.goal, ","))
	}

	for i := 0; i < size - 1; i++ {
		if machine.joltage {
			initial += "0,"
		} else {
			initial += "."
		}
	}

	if machine.joltage {
		initial += "0"
	} else {
		initial += "."
	}

	return initial
}

func dp(machine Machine, state string, buttons []Button, seen map[string]struct{}, pressed []Button, bestSoFar int, minimals map[string]int) int {

	if state == machine.goal {
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
		newLights := changeState(machine, state, button.changes)

		copied := copyList(pressed)
		copied = append(copied, button)
		cost := dp(machine, newLights, buttons, copyMap(seen), copied, answer, minimals) + 1
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

func changeState(machine Machine, state string, changes []int) string {

	if !machine.joltage {
		runes := []rune(state)
		for _, change := range changes {
			if runes[change] == '.' {
				runes[change] = '#'
			} else {
				runes[change] = '.'
			}
		}

		changed := string(runes)
		return changed
	} else {
		parts := strings.Split(state, ",")
		for _, change := range changes {
			value, _ := strconv.Atoi(parts[change])
			value++
			parts[change] = strconv.Itoa(value)
		}

		return strings.Join(parts, ",")
	}
}

func buildMachine(line string, useJoltage bool) Machine {
	parts := strings.Split(line, " ")

	var goal string
	if !useJoltage {
		goal = parts[0]
	} else {
		goal = parts[len(parts)-1]
	}

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

	return Machine{goal[1 : len(goal)-1], buttons, useJoltage}
}

func copyMap(original map[string]struct{}) map[string]struct{} {
	copiedMap := make(map[string]struct{}, len(original))
	for key, value := range original {
		copiedMap[key] = value
	}

	return copiedMap
}
