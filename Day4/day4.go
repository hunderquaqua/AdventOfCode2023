package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")
	var sum int
	instances := make(map[int]int)
	total := 0
	for i := 0; i < len(data); i++ {
		instances[i] = 1
	}

	for i, line := range data {
		// Splitting the line and storing the number before and after '|'
		line = line[10:]
		splittedLine := strings.Split(line, "|")
		myNums := getStrings(splittedLine[1])
		win := getStrings(splittedLine[0])
		var match int
		// Incrementing match if the it's a winning number
		for _, num := range myNums {
			if contains(win, num) {
				match++
			}
		}

		// Part 1 - Finding the sum of all matches
		if match > 0 {
			sum += int(math.Pow(2, float64(match-1)))
		} else {
			continue
		}

		// Part 2
		if i < len(data) {
			for j := 1; j <= match; j++ {
				instances[i+j] += instances[i]
			}
		}

	}
	for i := 0; i < len(instances); i++ {
		total += instances[i]
	}
	fmt.Println(sum)
	fmt.Println(total)
}

func getStrings(card string) []string {
	r, _ := regexp.Compile(`\d+`)
	return r.FindAllString(card, -1)
}

func contains(win []string, num string) bool {
	for _, el := range win {
		if el == num {
			return true
		}
	}
	return false
}
