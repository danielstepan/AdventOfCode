package solution

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielstepan/adventofcode/internal/input"
)

func prepareInput(inputFilePath string) ([]string, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	return content, nil
}

func canBeCalculated(numbers []int, target int, idx int, currentVal int, concatenate bool) bool {
	if idx == len(numbers) {
		return currentVal == target
	}
	if currentVal == 0 {
		currentVal = numbers[0]
	}

	if concatenate {
		concatenated_value, _ := strconv.Atoi(strconv.Itoa(currentVal) + strconv.Itoa(numbers[idx]))
		if canBeCalculated(numbers, target, idx+1, concatenated_value, concatenate) {
			return true
		}
	}

	return canBeCalculated(numbers, target, idx+1, currentVal+numbers[idx], concatenate) || canBeCalculated(numbers, target, idx+1, currentVal*numbers[idx], concatenate)
}

func run(inputFilePath string, concatenate bool) (string, error) {
	content, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	sum := 0

	for _, line := range content {
		splittedNumbers := strings.Split(line, ":")

		result, _ := strconv.Atoi(splittedNumbers[0])
		numbers, _ := input.StringSliceToIntSlice(strings.Fields(splittedNumbers[1]))

		if canBeCalculated(numbers, result, 1, 0, concatenate) {
			sum += result
		}
	}

	return strconv.Itoa(sum), nil
}

func Part1(inputFilePath string) (string, error) {
	return run(inputFilePath, false)
}

func Part2(inputFilePath string) (string, error) {
	return run(inputFilePath, true)
}
