package similarity

import (
	"errors"

	"github.com/danielstepan/adventofcode/internal/util"
)

func GetSimilarity(list1, list2 []int) (int, error) {
	if !util.AreIntSlicesEqualLength(list1, list2) {
		return 0, errors.New("slices must be of equal length")
	}

	similarity := 0
	for _, num := range list1 {
		numCount := 0
		for _, num2 := range list2 {
			if num == num2 {
				numCount++
			}
		}
		similarity += num * numCount
	}
	return similarity, nil
}
