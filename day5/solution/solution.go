package solution

import (
	"slices"
	"strconv"
	"strings"

	"github.com/danielstepan/adventofcode/internal/input"
)

type Rule struct {
	Earlier int
	Later   int
}

func parseRule(r []string) (Rule, error) {
	firstNumber, err := strconv.Atoi(r[0])
	if err != nil {
		return Rule{}, err
	}
	secondNumber, err := strconv.Atoi(r[1])
	if err != nil {
		return Rule{}, err
	}

	return Rule{
		Earlier: firstNumber,
		Later:   secondNumber,
	}, nil
}

func parseNumbers(line string) ([]int, error) {
	numbers := []int{}
	for _, n := range strings.Split(line, ",") {
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func prepareInput(inputFilePath string) ([]Rule, [][]int, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return nil, nil, err
	}

	rules := []Rule{}
	manualNumbers := [][]int{}

	for _, line := range content {
		if line == "" {
			continue
		}

		r := strings.Split(line, "|")
		if len(r) == 2 {
			rule, err := parseRule(r)
			if err != nil {
				return nil, nil, err
			}
			rules = append(rules, rule)
		} else {
			numbers, err := parseNumbers(line)
			if err != nil {
				return nil, nil, err
			}
			manualNumbers = append(manualNumbers, numbers)
		}
	}

	return rules, manualNumbers, nil
}

func Part1(inputFilePath string) (string, error) {
	rules, manualNumbers, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	sum := 0

	for _, numbers := range manualNumbers {
		if !isOrderBroken(numbers, rules) {
			sum += numbers[len(numbers)/2]
		}
	}

	return strconv.Itoa(sum), nil
}

func Part2(inputFilePath string) (string, error) {
	rules, manualNumbers, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	sum := 0
	brokenLines := [][]int{}

	for _, numbers := range manualNumbers {
		if isOrderBroken(numbers, rules) {
			brokenLines = append(brokenLines, numbers)
		}
	}

	for _, numbers := range brokenLines {
		fixBrokenOrder(numbers, rules)
		sum += numbers[len(numbers)/2]
	}
	return strconv.Itoa(sum), nil
}

func fixBrokenOrder(numbers []int, rules []Rule) {
	for {
		isBroken := false
		for _, r := range rules {
			earlierIdx := slices.Index(numbers, r.Earlier)
			laterIdx := slices.Index(numbers, r.Later)

			if earlierIdx == -1 || laterIdx == -1 {
				continue
			}

			if earlierIdx > laterIdx {
				isBroken = true
				numbers[earlierIdx], numbers[laterIdx] = numbers[laterIdx], numbers[earlierIdx]
			}
		}
		if !isBroken {
			break
		}
	}
}

func isOrderBroken(numbers []int, rules []Rule) bool {
	for _, r := range rules {
		earlierIdx := slices.Index(numbers, r.Earlier)
		laterIdx := slices.Index(numbers, r.Later)

		if earlierIdx == -1 || laterIdx == -1 {
			continue
		}

		if earlierIdx > laterIdx {
			return true
		}
	}
	return false
}
