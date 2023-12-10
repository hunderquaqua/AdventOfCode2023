package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := part1(file)
	ans := part2(file)
	fmt.Println(sum)
	fmt.Println(ans)

}

func part1(file []byte) int {
	grid := strings.Split(string(file), "\n")
	cs := make(map[string]bool)

	// Go through each element in grid
	for r, row := range grid {
		for c, ch := range row {
			// Skipping if ch is not a symbol
			if unicode.IsDigit(ch) || ch == '.' || ch == 13 {
				continue
			}
			// Checking if the symbol is surrounded by a digit
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					newR, newC := r+dr, c+dc
					if newR < 0 || newR >= len(grid) || newC < 0 || newC >= len(grid[newR]) || !unicode.IsDigit(rune(grid[newR][newC])) {
						continue
					}
					// Getting the whole number, going back one element at a time
					for newC > 0 && unicode.IsDigit(rune(grid[newR][newC-1])) {
						newC--
					}
					// Adding the coordinates of the number around the symbol into the map as a key
					cs[fmt.Sprintf("%d,%d", newR, newC)] = true
				}
			}
		}
	}
	// An array to store all the values of the numbers which coordinates we store
	ns := make([]int, 0)

	for coord := range cs {
		r, _ := strconv.Atoi(strings.Split(coord, ",")[0])
		c, _ := strconv.Atoi(strings.Split(coord, ",")[1])
		str := ""
		for c < len(grid[r]) && unicode.IsDigit(rune(grid[r][c])) {
			str += string(grid[r][c])
			c++
		}
		num, _ := strconv.Atoi(str)
		ns = append(ns, num)
	}

	sum := 0
	for _, num := range ns {
		sum += num
	}
	return sum
}

func part2(file []byte) int {
	grid := strings.Split(string(file), "\n")
	total := 0
	// Go through each element in grid
	for r, row := range grid {
		for c, ch := range row {
			// Skipping if ch is not a symbol
			if ch != '*' {
				continue
			}
			cs := make(map[string]bool)
			// Checking if the symbol is surrounded by a digit
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					newR, newC := r+dr, c+dc
					if newR < 0 || newR >= len(grid) || newC < 0 || newC >= len(grid[newR]) || !unicode.IsDigit(rune(grid[newR][newC])) {
						continue
					}
					// Getting the whole number, going back one element at a time
					for newC > 0 && unicode.IsDigit(rune(grid[newR][newC-1])) {
						newC--
					}
					// Adding the coordinates of the number around the symbol into the map as a key
					cs[fmt.Sprintf("%d,%d", newR, newC)] = true
				}
			}
			ns := make([]int, 0)

			for coord := range cs {
				r, _ := strconv.Atoi(strings.Split(coord, ",")[0])
				c, _ := strconv.Atoi(strings.Split(coord, ",")[1])
				str := ""
				for c < len(grid[r]) && unicode.IsDigit(rune(grid[r][c])) {
					str += string(grid[r][c])
					c++
				}
				num, _ := strconv.Atoi(str)
				ns = append(ns, num)
			}
			if len(ns) == 2 {
				total += ns[0] * ns[1]
			}
		}
	}
	return total
}
