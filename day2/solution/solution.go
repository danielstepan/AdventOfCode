package solution

import (
	"math"
	"slices"
	"sort"
	"strconv"

	"github.com/danielstepan/adventofcode/day2/internal/parser"
	"github.com/danielstepan/adventofcode/internal/input"
)

func prepareInput(inputFilePath string) ([][]int, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return [][]int{}, err
	}

	sliceOfSlices, err := parser.SplitStringsToSlicesOfInts(content)
	if err != nil {
		return [][]int{}, err
	}

	return sliceOfSlices, nil
}

func Part1(inputFilePath string) (string, error) {
	sliceOfSlices, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	saveSum := 0

	for _, slice := range sliceOfSlices {
		if isSliceSafe(slice) {
			saveSum++
		}
	}

	return strconv.Itoa(saveSum), nil
}

func Part2(inputFilePath string) (string, error) {
	sliceOfSlices, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	saveSum := 0

	for _, slice := range sliceOfSlices {
		if isSliceSafe(slice) {
			saveSum++
			continue
		}

		for i := 0; i < len(slice); i++ {
			modifiedSlice := removeIndex(slice, i)
			if isSliceSafe(modifiedSlice) {
				saveSum++
				break
			}
		}
	}

	return strconv.Itoa(saveSum), nil
}

func isSliceSafe(slice []int) bool {
	reversed := make([]int, len(slice))
	copy(reversed, slice)
	slices.Reverse(reversed)

	sorted := sort.IntsAreSorted(slice) || sort.IntsAreSorted(reversed)
	diffInLimit := true

	for i := 0; i < len(slice)-1; i++ {
		diff := int(math.Abs(float64(slice[i+1]) - float64(slice[i])))
		if diff == 0 || diff > 3 {
			diffInLimit = false
			break
		}
	}

	return sorted && diffInLimit
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
