package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")

	// Part 1
	times := extractNumbers(data[0])
	distances := extractNumbers(data[1])

	n := 1

	for i, time := range times {
		distance := distances[i]
		margin := 0
		for hold := 0; hold < time; hold++ {
			if hold*(time-hold) > distance {
				margin++
			}
		}
		n *= margin
	}

	fmt.Println(n)
	// Part 2
	n = 0
	time := extractNumber(data[0])
	distance := extractNumber(data[1])

	for hold := 0; hold < time; hold++ {
		if hold*(time-hold) > distance {
			n++
		}
	}
	fmt.Println(n)
}

func extractNumbers(str string) []int {
	r, _ := regexp.Compile(`\d+`)
	strings := r.FindAllString(str, -1)
	var nums []int

	for _, strNum := range strings {
		intNum, _ := strconv.Atoi(strNum)
		nums = append(nums, intNum)
	}
	return nums
}

func extractNumber(str string) int {
	r, _ := regexp.Compile(`\d+`)
	strings := r.FindAllString(str, -1)
	var numStr string
	for _, num := range strings {
		numStr += num
	}
	num, _ := strconv.Atoi(numStr)
	return num
}
