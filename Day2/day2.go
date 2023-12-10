package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(file)
	values := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var sum, sumMins int
	for id, line := range strings.Split(text, "\n") {

		// Deleting the "Game _: " prefix from every line
		prefix := "Game " + strconv.Itoa(id+1) + ": "
		line = strings.TrimSpace(strings.TrimPrefix(line, prefix))

		// Flag to check if no value exceeds the corresponding value in the map
		noColorExceeds := true

		// Setting minimal possible values to play the game
		var red, green, blue int

		// Going through each round of every game
		for _, round := range strings.Split(line, ";") {

			// Going through each value-color pair in rounds
			for _, instance := range strings.Split(round, ",") {
				instance = strings.TrimSpace(instance)
				instanceParsed := strings.Fields(instance)
				value, _ := strconv.Atoi(instanceParsed[0])

				// Checking if the value-color is less than the one we set for the color
				if values[instanceParsed[1]] < value {
					noColorExceeds = false
				}

				// Assigning the colors minimal count of balls to game
				if instanceParsed[1] == "red" {
					red = max(red, value)
				} else if instanceParsed[1] == "green" {
					green = max(green, value)
				} else {
					blue = max(blue, value)
				}
			}
		}
		if noColorExceeds {
			sum += (id + 1)
		}
		sumMins += red * green * blue
	}

	fmt.Println("Sum:", sum)
	fmt.Println("SumMin: ", sumMins)
}
