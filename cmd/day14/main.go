package main

import (
	"fmt"

	"github.com/danielstepan/adventofcode/day14/solution"
)

func main() {
	result1, err := solution.Part1("day14/solution/input.txt", 101, 103, 100)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Day 14, Part 1:", result1)

	result, err := solution.Part2("day14/solution/input.txt", 101, 103)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Day 14, Part 2:", result)
}
