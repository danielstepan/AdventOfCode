package parser

import (
	"strconv"
	"strings"
)

func SplitStringsToSlicesOfInts(content []string) (sliceOfSlices [][]int, err error) {
	for _, line := range content {
		lineElements := strings.Fields(line)
		var s []int
		for _, element := range lineElements {
			val, err := strconv.Atoi(element)
			if err != nil {
				return [][]int{}, err
			}
			s = append(s, val)
		}
		sliceOfSlices = append(sliceOfSlices, s)
	}
	return sliceOfSlices, nil
}
