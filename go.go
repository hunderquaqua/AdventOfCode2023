package main

import "fmt"

func main() {
	steps := "abcdefg"

	for i := 0; i < 10; i++ {
		steps = steps[1:] + string(steps[0])
		fmt.Println(steps)
	}

}
