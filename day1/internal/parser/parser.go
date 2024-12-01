package parser

import (
	"strconv"
	"strings"
)

func SplitStringsToInts(content []string) (s1, s2 []int, err error) {
	for _, line := range content {
		s := strings.Fields(line)
		firstVal, err := strconv.Atoi(s[0])
		if err != nil {
			return []int{}, []int{}, err
		}
		secondVal, err := strconv.Atoi(s[1])
		if err != nil {
			return []int{}, []int{}, err
		}
		s1 = append(s1, firstVal)
		s2 = append(s2, secondVal)
	}
	return s1, s2, nil
}
