package main

import (
	"fmt"

	"github.com/danielstepan/adventofcode/day4/solution"
)

func main() {
	result1, err := solution.Part1("day4/solution/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Day 4, Part 1:", result1)

	result2, err := solution.Part2("day4/solution/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Day 4, Part 2:", result2)
}
