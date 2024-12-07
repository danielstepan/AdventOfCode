package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileLines(name string) (content []string, err error) {
	file, err := os.Open(name)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content, nil
}

func StringSliceToIntSlice(stringSlice []string) ([]int, error) {
	intSlice := make([]int, len(stringSlice))
	for i, str := range stringSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("error converting %s: %v", str, err)
		}
		intSlice[i] = num
	}
	return intSlice, nil
}
