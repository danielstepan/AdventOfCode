package solution

import (
	"fmt"
	"strconv"

	"github.com/danielstepan/adventofcode/internal/input"
)

func prepareInput(inputFilePath string) ([]int, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	intSlice, err := input.StringToIntSlice(content[0])
	return intSlice, nil
}

func Part1(inputFilePath string) (string, error) {
	return run(inputFilePath, 25)
}

func Part2(inputFilePath string) (string, error) {
	stones, err := prepareInput(inputFilePath)
	if err != nil {
		return "", fmt.Errorf("could not prepare input: %w", err)
	}

	stonesDict := map[int]int{}
	for _, num := range stones {
		stonesDict[num]++
	}

	for i := 0; i < 75; i++ {
		tmpStonesDict := map[int]int{}

		for num, count := range stonesDict {
			if num == 0 {
				tmpStonesDict[1] += count
			} else if digitCount(num)%2 == 0 {
				numStr := strconv.Itoa(num)
				half := len(numStr) / 2
				left, _ := strconv.Atoi(numStr[:half])
				right, _ := strconv.Atoi(numStr[half:])

				tmpStonesDict[left] += count
				tmpStonesDict[right] += count
			} else {
				tmpStonesDict[num*2024] += count
			}
		}

		stonesDict = tmpStonesDict
	}

	count := 0
	for _, v := range stonesDict {
		count += v
	}

	return strconv.Itoa(count), nil
}

func run(inputFilePath string, blinks int) (string, error) {
	stones, err := prepareInput(inputFilePath)
	if err != nil {
		return "", fmt.Errorf("could not prepare input: %w", err)
	}

	for i := 0; i < blinks; i++ {
		tmpStones := []int{}
		for _, num := range stones {
			if num == 0 {
				tmpStones = append(tmpStones, 1)
			} else if digitCount(num)%2 == 0 {
				numStr := strconv.Itoa(num)
				half := len(numStr) / 2
				left, _ := strconv.Atoi(numStr[:half])
				right, _ := strconv.Atoi(numStr[half:])
				tmpStones = append(tmpStones, left)
				tmpStones = append(tmpStones, right)
			} else {
				tmpStones = append(tmpStones, num*2024)
			}
		}

		stones = tmpStones
	}

	count := len(stones)

	return strconv.Itoa(count), nil
}

func digitCount(num int) int {
	return len(strconv.Itoa(num))
}
