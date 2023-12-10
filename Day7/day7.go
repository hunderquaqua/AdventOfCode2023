package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CustomStringSort []string

var customOrder = "23456789TJQKA"

func (c CustomStringSort) Len() int { return len(c) }
func (c CustomStringSort) Less(i, j int) bool {
	return customHandLess(c[i], c[j])
}
func (c CustomStringSort) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func parseHand(hand string) []byte {
	return []byte(strings.TrimSpace(hand))
}
func customHandLess(a, b string) bool {
	orderMap := make(map[byte]int)
	for i, ch := range customOrder {
		orderMap[byte(ch)] = i
	}

	// Parsing the hands
	handA := parseHand(a)
	handB := parseHand(b)

	// Comparing the hands based on type
	typeA, typeB := getHandType(handA), getHandType(handB)
	if typeA != typeB {
		return typeA < typeB
	}

	// If hands are the same type, we sort them by cards
	for i := 0; i < 5; i++ {
		if orderMap[handA[i]] != orderMap[handB[i]] {
			return orderMap[handA[i]] > orderMap[handB[i]]
		}
	}
	return false
}

func getHandType(hand []byte) string {
	counts := make(map[byte]int)
	for _, card := range hand {
		counts[card]++
	}

	for card, count := range counts {
		switch count {
		case 5:
			return "FiveOfAKind"
		case 4:
			return "FourOfAKind"
		case 3:
			for _, otherCount := range counts {
				if otherCount == 2 {
					return "FullHouse"
				}
			}
			return "ThreeOfAKind"
		case 2:
			for otherCard, otherCount := range counts {
				if otherCount == 2 && card != otherCard {
					return "TwoPair"
				}
			}
			return "OnePair"
		}
	}

	return "HighCard"
}

func main() {
	file, _ := os.ReadFile("input.txt")

	games := strings.Split(string(file), "\n")
	gameMap := make(map[string]int)
	sum := 0
	for _, game := range games {
		splittedLine := strings.Split(game, " ")
		num := atoi(splittedLine[1])
		gameMap[splittedLine[0]] = num
	}

	var slice []string
	for k := range gameMap {
		slice = append(slice, k)
	}

	sort.Sort(CustomStringSort(slice))
	for i, v := range slice {
		sum += (i + 1) * gameMap[v]
		fmt.Println(i+1, gameMap[v])
	}
	fmt.Println(gameMap)
	fmt.Println(sum)

}

func atoi(str string) int {
	str = strings.TrimSpace(str)
	num, _ := strconv.Atoi(str)
	return num
}
