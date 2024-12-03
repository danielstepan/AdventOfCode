package solution

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/danielstepan/adventofcode/internal/input"
)

func Part1(inputFilePath string) (string, error) {
	return run(inputFilePath, `(mul)\((\d{1,3}),(\d{1,3})\)`)
}

func Part2(inputFilePath string) (string, error) {
	return run(inputFilePath, `do\(\)|don't\(\)|(mul)\((\d{1,3}),(\d{1,3})\)`)
}

func prepareInput(inputFilePath string) ([]string, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return []string{""}, err
	}

	return content, nil
}

func run(inputFilePath string, regex string) (string, error) {
	content, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	r, _ := regexp.Compile(regex)

	isEvaluationEnabled := true
	sum := 0

	for _, line := range content {
		match := r.FindAllStringSubmatch(line, -1)
		for _, m := range match {
			result, err := evaluateExpression(m, &isEvaluationEnabled)
			if err != nil {
				return "", fmt.Errorf("error evaluating expression: %v", err)
			}
			sum += result
		}
	}

	return strconv.Itoa(sum), nil
}

func evaluateExpression(expression []string, isEvaluationEnabled *bool) (int, error) {
	if expression[0] == "do()" {
		*isEvaluationEnabled = true
		return 0, nil
	} else if expression[0] == "don't()" {
		*isEvaluationEnabled = false
		return 0, nil
	}

	if *isEvaluationEnabled == false {
		return 0, nil
	}

	operation := expression[1]
	operand1, err := strconv.Atoi(expression[2])
	if err != nil {
		return 0, fmt.Errorf("error converting operand1 to int: %v", err)
	}
	operand2, err := strconv.Atoi(expression[3])
	if err != nil {
		return 0, fmt.Errorf("error converting operand2 to int: %v", err)
	}

	if operation == "mul" {
		return operand1 * operand2, nil
	}

	return 0, fmt.Errorf("unknown operation: %s", operation)
}
