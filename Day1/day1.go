package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const input1 = "input.txt"

func main() {
	// Opening the the txt file
	// Creating a scanner to read the file line by line
	file := OpenFile(input1)
	defer file.Close()

	result1, result2, err := SumOfValues(file)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("Sum of the Calibration Values: ", result1)
	fmt.Println("Sum of the Calibration Values including spelled numbers: ", result2)
}

func OpenFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error occured when opening file: ", err)
		os.Exit(1)
	}
	return file
}

func SumOfValues(file *os.File) (int, int, error) {
	scanner := bufio.NewScanner(file)
	sum := 0
	sum2 := 0
	// Iterating through the lines
	for scanner.Scan() {
		line := scanner.Text()
		sum += LineParser(line)
		sum2 += LineParserSpelled(line)
	}

	// Checking for errors while scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

	return sum, sum2, nil
}

func LineParser(line string) int {
	var first rune
	var last rune
	for _, r := range line {
		if unicode.IsDigit(r) && first == 0 {
			first = r
			last = r
			continue
		}
		if unicode.IsDigit(r) && last != 0 {
			last = r
		}
	}
	result, err := strconv.Atoi(string(first) + string(last))

	if err != nil {
		fmt.Println("Error while converting string to int -> ", err)
	}
	return result
}

func LineParserSpelled(line string) int {
	var first, last string
	nums := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"0":     "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	for i := 0; i < len(line); i++ {
		for j := i + 1; j <= len(line); j++ {
			substring := line[i:j]
			if num, exists := nums[substring]; exists {
				if first == "" {
					first = num
				}

				last = num
			}
		}

		char := string(line[i])
		if num, exists := nums[char]; exists {
			if first == "" {
				first = num
			}

			last = num
		}
	}

	result, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Println("Error converting string to int -> ", err)
		os.Exit(1)
	}
	return result
}
