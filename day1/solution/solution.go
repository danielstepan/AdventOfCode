package solution

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/danielstepan/adventofcode/day1/internal/distance"
	"github.com/danielstepan/adventofcode/day1/internal/parser"
	"github.com/danielstepan/adventofcode/day1/internal/similarity"
	"github.com/danielstepan/adventofcode/internal/input"
)

func prepareInput() ([]int, []int, error) {
	content, err := input.ReadFileLines("day1/solution/input.txt")
	if err != nil {
		return []int{}, []int{}, errors.New("Error reading input file")
	}

	list1, list2, err := parser.SplitStringsToInts(content)
	if err != nil {
		return []int{}, []int{}, errors.New("Error parsing input")
	}
	sort.Ints(list1)
	sort.Ints(list2)
	return list1, list2, nil
}

func Part1() (string, error) {
	list1, list2, err := prepareInput()
	if err != nil {
		return "", err
	}

	diff, err := distance.SumOfDifference(list1, list2)
	if err != nil {
		fmt.Errorf("Error calculating sum of difference: %v", err)
		return "", err
	}
	return strconv.Itoa(diff), nil
}

func Part2() (string, error) {
	list1, list2, err := prepareInput()

	similarity, err := similarity.GetSimilarity(list1, list2)
	if err != nil {
		fmt.Errorf("Error calculating similarity: %v", err)
		return "", err
	}
	return strconv.Itoa(similarity), nil
}
