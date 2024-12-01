package solution

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/danielstepan/adventofcode/day1/internal/distance"
	"github.com/danielstepan/adventofcode/day1/internal/parser"
	"github.com/danielstepan/adventofcode/day1/internal/similarity"
	"github.com/danielstepan/adventofcode/internal/input"
)

func prepareInput(inputFilePath string) ([]int, []int, error) {
	fmt.Println("Reading input file", inputFilePath)
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return []int{}, []int{}, err
	}

	list1, list2, err := parser.SplitStringsToInts(content)
	if err != nil {
		return []int{}, []int{}, err
	}
	sort.Ints(list1)
	sort.Ints(list2)
	return list1, list2, nil
}

func Part1(inputFilePath string) (string, error) {
	list1, list2, err := prepareInput(inputFilePath)
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

func Part2(inputFilePath string) (string, error) {
	list1, list2, err := prepareInput(inputFilePath)

	similarity, err := similarity.GetSimilarity(list1, list2)
	if err != nil {
		fmt.Errorf("Error calculating similarity: %v", err)
		return "", err
	}
	return strconv.Itoa(similarity), nil
}
