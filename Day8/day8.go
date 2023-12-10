package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("demo.txt")
	//ans1 := part1(file)
	//fmt.Println(ans1)
	part2(file)
}

func part1(file []byte) int {
	data := strings.Split(string(file), "\n")
	network := make(map[string][]string)
	steps := strings.TrimSpace(data[0])
	rest := data[2:]

	for _, line := range rest {
		parsedLine := strings.Split(line, " = ")
		network[parsedLine[0]] = parseString(parsedLine[1])
	}

	stepCount := 0
	current := "AAA"

	for current != "ZZZ" {
		stepCount++

		if steps[0] == 'L' {
			current = network[current][0]
		} else {
			current = network[current][1]
		}
		steps = steps[1:] + string(steps[0])
	}
	return stepCount
}

func part2(file []byte) {
	data := strings.Split(string(file), "\n")
	network := make(map[string][]string)
	steps := strings.TrimSpace(data[0])
	rest := data[2:]

	for _, line := range rest {
		parsedLine := strings.Split(line, " = ")
		network[parsedLine[0]] = parseString(parsedLine[1])
	}

	var positions []string
	for key := range network {
		if strings.HasSuffix(key, "A") {
			positions = append(positions, key)
		}
	}

	firstZ := ""
	var cycles [][]int
	for _, current := range positions {
		var cycle []int

		currentSteps := steps
		stepCount := 0

		for {
			for stepCount == 0 || !strings.HasSuffix(current, "Z") {
				stepCount++
				if steps[0] == 'L' {
					current = network[current][0]
				} else {
					current = network[current][1]
				}
				currentSteps = currentSteps[1:] + string(currentSteps[0])
			}
			cycle = append(cycle, stepCount)
			if firstZ == "" {
				firstZ = current
				stepCount = 0
			} else if current == firstZ {
				break
			}
			cycles = append(cycles, cycle)
		}
	}

	fmt.Println(cycles)
}

func parseString(input string) []string {
	// Trim the leading "(" and trailing ")" from the string
	trimmed := strings.TrimSpace(input)
	trimmed = strings.TrimLeft(trimmed, "(")
	trimmed = strings.TrimRight(trimmed, ")")

	// Split the string by ", "
	elements := strings.Split(trimmed, ", ")

	return elements
}
